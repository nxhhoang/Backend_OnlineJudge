import crypto from 'crypto'
import nodemailer from 'nodemailer'
import passport from 'passport'
import { User, IUser } from '../mongodb/user'
import { Request, Response, NextFunction } from 'express'
import { IVerifyOptions } from 'passport-local'
import { body, validationResult } from 'express-validator'
import '../config/passport'
require("dotenv").config();

/**
  * Login @Route POST /auth/login
*/
export const postLogin = async (req: Request, res: Response, next: NextFunction) => {
  await body('password')
    .isLength({ min: 8 }).withMessage('Password must be at least 8 characters')
    .run(req);
  await body('username')
    .notEmpty().withMessage('Username cannot be empty')
    .isAlphanumeric().withMessage('Username must be alphanumeric')
    .isLength({ min: 6 }).withMessage('Username must be at least 6 characters')
    .run(req);

  const errors = validationResult(req);

  if (!errors.isEmpty()) {
    res.status(400).json({ errors: errors.array() });
    return;
  }

  passport.authenticate('local', (err: Error, user: IUser, info: IVerifyOptions) => {
    if (err) {
      res.status(500).json({ message: 'An error occurred during authentication' });
      return;
    }
    if (!user) {
      res.status(400).json({ message: info.message || 'Invalid username or password' });
      return;
    }

    req.logIn(user, (loginErr) => {
      if (loginErr) {
        res.status(500).json({ message: 'Could not login' });
        return;
      }
      res.status(200).json({
        message: 'Login successfully',
        user: {
          id: user.id,
          username: user.username,
          email: user.email
        }
      });
      return;
    });
  })(req, res, next);
};

/**
  * Logout @Route GET /logout
*/
export const getLogout = (req: Request, res: Response): void => {
  req.logout((err) => {
    if (err) {
      res.status(500).json({ message: 'Error during logging out' });
      return;
    }
    res.status(200).json({ message: 'Log out successfully' });
  });
}

/**
  * SignUp @Route POST /auth/signup
*/
export const postSignup = async (req: Request, res: Response, next: NextFunction) => {
  await body('email')
    .not().isEmpty().withMessage('Email must not be empty')
    .isEmail().withMessage('Email is Invalid')
    .run(req);
  await body('password')
    .isLength({ min: 8 }).withMessage('Password must be at least 8 characters')
    .run(req);
  await body('username')
    .notEmpty().withMessage('Username cannot be empty')
    .isAlphanumeric().withMessage('Username must be alphanumeric')
    .isLength({ min: 6 }).withMessage('Username must be at least 6 characters')
    .run(req);

  const errors = validationResult(req);

  if (!errors.isEmpty()) {
    res.status(500).json({ errors: errors.array() });
    return;
  }

  const user = new User({
    email: req.body.email,
    password: req.body.password,
    username: req.body.username
  })

  try {
    if (await User.findOne({ email: req.body.email })) {
      res.status(400).json({ message: 'A user with that email already exists.' });
      return;
    }

    if (await User.findOne({ username: req.body.username })) {
      res.status(400).json({ message: 'A user with that username already exists.' });
      return;
    }
    const savedUser = await user.save();
    req.logIn(savedUser, (loginErr) => {
      if (loginErr) {
        res.status(500).json({ message: 'Could not login' });
        return;
      }
      res.status(200).json({
        message: 'Sign up successfully',
        user: {
          id: savedUser.id,
          username: savedUser.username,
          email: savedUser.email
        }
      });
    });
  } catch (err) {
    console.log(err);
    res.status(500).json({ message: 'An error occurred during sign up' });
  }
}

/**
  * @Route POST /password
*/
export const postPassword = async (req: Request, res: Response, next: NextFunction) => {
  await body('password')
    .isLength({ min: 8 }).withMessage('Password must be at least 8 characters')
    .run(req);
  await body('confirm')
    .equals(req.body.password).withMessage('Passwords do not match')
    .run(req);

  const errors = validationResult(req);

  if (!errors.isEmpty()) {
    res.status(500).json({ errors: errors.array() });
    return;
  }

  const user = req.user as IUser;
  try {
    const fUser = await User.findById(user.id);
    if (!fUser) {
      res.status(404).json({ message: 'User not found' });
      return;
    } else {
      fUser.password = req.body.password;
      const savedUser = await fUser.save();
      res.status(200).json({
        message: 'Password has been changed successfully',
        user: {
          id: savedUser.id,
          username: savedUser.username,
          email: savedUser.email
        }
      });
    }
  } catch (err) {
    res.status(500).json({ message: 'An error occurred when attempt to change password' });
    return;
  }
};

/**
  * @Route POST /auth/forgot
*/
export const postForgot = async (req: Request, res: Response, next: NextFunction) => {
  await body('email')
    .isEmail().withMessage('Please enter a valid email')
    .normalizeEmail({ gmail_remove_dots: false })
    .run(req);

  const errors = validationResult(req);

  if (!errors.isEmpty()) {
    res.status(500).json({ errors: errors.array() });
    return;
  }
  try {
    const token = crypto.randomBytes(16).toString('hex');
    const user = await User.findOne({ email: req.body.email });
    if (!user) {
      res.status(400).json({ message: 'No account with the give email' });
      return;
    }
    user.passwordResetToken = token;
    user.passwordResetTokenExpires = new Date(new Date().getTime() + 60 * 60 * 1000); // 1 hour
    await user.save();

    const transporter = nodemailer.createTransport({
      service: 'gmail',
      auth: {
        user: process.env.AUTH_GOOGLE_GMAIL,
        pass: process.env.AUTH_GOOGLE_PASSWORD
      }
    });

    const mailOptions = {
      to: user.email,
      from: 'online-judge-recover@gmail.com',
      subject: 'Reset your password on Online-judge',
      text: `This email is sent because you (or someone else) have requested to change your password of your account.
      \n\nPlease click on the following link, or paste into your browser to continue the process:\n\n
    http://${req.headers.host}/auth/reset/${token}\n\nIf you don't want to change your password, please ignore this email.`
    };
    const info = await transporter.sendMail(mailOptions);
    res.status(200).json({ message: `An email has been sent to ${user.email}.` });
  } catch (err) {
    res.status(500).json({ message: 'An error occurred' });
  }
}

/**
  * @Route POST /reset/:token
*/
export const postReset = async (req: Request, res: Response, next: NextFunction) => {
  await body('password')
    .isLength({ min: 8 }).withMessage('Password must be at least 8 characters')
    .run(req);
  await body('confirm')
    .equals(req.body.password).withMessage('Passwords do not match')
    .run(req);

  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    res.status(500).json({ errors: errors.array() });
    return;
  }
  try {
    const user = await User.findOne({ passwordResetToken: req.params.token })
      .where('passwordResetTokenExpires').gt(Date.now());
    if (!user) {
      res.status(400).json({ message: 'Password reset token is invalid or has been expired' });
      return;
    }
    user.password = req.body.password;
    user.passwordResetToken = null;
    user.passwordResetTokenExpires = null;
    const savedUser = await user.save();
    req.logIn(savedUser, (loginErr) => {
      if (loginErr) {
        res.status(500).json({ message: 'Could not login' });
        return;
      }
      res.status(200).json({
        message: 'Password has been changed successfully',
        user: {
          id: savedUser.id,
          username: savedUser.username,
          email: savedUser.email
        }
      });
    });
  } catch (err) {
    res.status(500).json({ message: 'An error occurred' });
  }
}

/**
 * @route   GET /auth/status
 */
export const getAuthStatus = (req: Request, res: Response): void => {
  if (req.isAuthenticated()) {
    const user = req.user as IUser;
    res.status(200).json({
      isAuthenticated: true,
      user: {
        id: user.id,
        username: user.username,
        email: user.email,
      }
    });
  } else {
    res.status(401).json({
      isAuthenticated: false,
      user: null
    });
  }
};

/**
  * @Route    GET /auth/csrftoken
*/
export const getCsrfToken = (req: Request, res: Response) => {
  if (res.locals._csrf) {
    res.status(200).json({ csrfToken: res.locals._csrf, });
    return;
  }
  res.status(400).json({ message: 'No csrf token' });
}
