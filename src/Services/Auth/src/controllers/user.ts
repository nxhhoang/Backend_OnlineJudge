import crypto from 'crypto'
import nodemailer from 'nodemailer'
import passport from 'passport'
import { User, IUser } from '../mongodb/user'
import { Request, Response, NextFunction } from 'express'
import { IVerifyOptions } from 'passport-local'
import { Result, ValidationError, body, check, validationResult } from 'express-validator'
import '../config/passport'
require("dotenv").config();

/**
  * Login @Route GET /login
*/
export const getLogin = (req: Request, res: Response): void => {
  if (req.user) {
    return res.redirect('/');
  }
  const token = res.locals._csrf;
  const errorMessages = req.flash('errors') as unknown as string[];
  const passportErrors = req.flash('error') as unknown as string[];
  const allErrors = [...errorMessages, ...passportErrors];

  res.render('login_signup', {
    title: 'Log in',
    token,
    allErrors,
    activeTab: 'login'
  });
}

/**
  * Login @Route POST /login
*/
export const postLogin = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
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
    req.flash("errors", errors.array().map(err => err.msg));
    return res.redirect('/auth/login');
  }

  passport.authenticate('local', (err: Error, user: IUser, info: IVerifyOptions) => {
    if (err) {
      req.flash('error', 'An error occurred during authentication');
      return res.redirect('/auth/login');
    }
    if (!user) {
      req.flash('error', info.message || 'Invalid username or password');
      return res.redirect('/auth/login');
    }

    req.logIn(user, (loginErr) => {
      if (loginErr) {
        req.flash('error', 'Could not login');
        return res.redirect('/auth/login');
      }
      return res.redirect('/');
    });
  })(req, res, next);
};

/**
  * Logout @Route GET /logout
*/
export const getLogout = (req: Request, res: Response): void => {
  req.logout((err) => {
    res.redirect('/auth/login');
  });
}

/**
  * SignUp @Route GET /signup
*/
export const getSignup = (req: Request, res: Response): void => {
  if (req.user) {
    return res.redirect("/");
  }

  const token = res.locals._csrf;
  const errorMessages = req.flash('errors') as unknown as string[];
  const passportErrors = req.flash('error') as unknown as string[];
  const allErrors = [...errorMessages, ...passportErrors];

  res.render('login_signup', {
    title: 'Sign up',
    token,
    allErrors,
    activeTab: 'signup'
  })
}

/**
  * SignUp @Route POST /signup
*/
export const postSignup = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
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
    req.flash("errors", errors.array().map(err => err.msg));
    return res.redirect('/auth/login');
  }

  const user = new User({
    email: req.body.email,
    password: req.body.password,
    username: req.body.username
  })

  try {
    if (await User.findOne({ email: req.body.email })) {
      req.flash('errors', 'A user with that email already exists.');
      return res.redirect('/auth/signup');
    }

    if (await User.findOne({ username: req.body.username })) {
      req.flash('errors', 'A user with that username already exists.');
      return res.redirect('/auth/signup');
    }
    const savedUser = await user.save();
    req.logIn(savedUser, (loginErr) => {
      if (loginErr) {
        req.flash('error', 'Could not login');
        return res.redirect('/auth/login');
      }
      return res.redirect('/');
    });

  } catch (err) {
    req.flash('errors', 'An error occurred');
    return res.redirect('/auth/signup');
  }
}

/**
  * @Route GET /profile
*/
export const getProfile = (req: Request, res: Response) => {
  if (req.isAuthenticated()) {
    // res.send(`<h1>Hello ${req.user}</h1><a href="/logout">Logout</a>`);
    const token = res.locals._csrf;
    res.render('profile', {
      titie: 'Profile',
      user: req.user as IUser,
      token,
      csrfToken: res.locals._csrf
    })
  } else {
    res.redirect('/auth/login');
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
    req.flash("errors", errors.array().map(err => err.msg));
    return res.redirect('/profile');
  }

  const user = req.user as IUser;
  try {
    const fUser = await User.findById(user.id);
    if (!fUser) {
      next();
    } else {
      fUser.password = req.body.password;
      const savedUser = await fUser.save();
      req.flash('success', 'Password has been changed');
      res.redirect('/profile');
    }
  } catch (err) {
    next(err);
  }
};

/**
  * @Route GET /forgot
*/
export const getForgot = (req: Request, res: Response): void => {
  if (req.isAuthenticated()) {
    return res.redirect('/');
  }

  res.render('forgot', {
    title: 'Forgot password',
    success: req.flash('success'),
    errors: req.flash('errors'),
    csrfToken: res.locals._csrf
  });
}

export const postForgot = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
  await body('email')
    .isEmail().withMessage('Please enter a valid email')
    .normalizeEmail({ gmail_remove_dots: false })
    .run(req);

  const errors = validationResult(req);

  if (!errors.isEmpty()) {
    req.flash("errors", errors.array().map(err => err.msg));
    return res.redirect('/auth/forgot');
  }
  try {
    const token = crypto.randomBytes(16).toString('hex');
    const user = await User.findOne({ email: req.body.email });
    if (!user) {
      req.flash('errors', 'No account with the given email');
      return res.redirect('/auth/forgot');
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
    console.log(info);
    req.flash('success', `An email has been sent to ${user.email}.`);
    return res.redirect('/auth/forgot');
  } catch (err) {
    console.log(err);
    req.flash('errors', 'An error occurred');
    return res.redirect('/auth/forgot');
  }
}

/**
  * @Route GET /reset
*/
export const getReset = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
  if (req.isAuthenticated()) {
    return res.redirect('/');
  }
  try {
    const user = await User.findOne({ passwordResetToken: req.params.token })
      .where('passwordResetTokenExpires').gt(Date.now());
    if (!user) {
      req.flash('errors', 'Password reset token is invalid or has been expired');
      return res.redirect('/auth/forgot');
    }
    res.render('reset', {
      title: 'Reset password',
      success: req.flash('success'),
      errors: req.flash('errors'),
      csrfToken: res.locals._csrf,
      resetToken: req.params.token
    });
  } catch (err) {
    next(err);
    return res.redirect('/auth/forgot');
  }
}

/**
  * @Route POST /reset
*/
export const postReset = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
  await body('password')
    .isLength({ min: 8 }).withMessage('Password must be at least 8 characters')
    .run(req);
  await body('confirm')
    .equals(req.body.password).withMessage('Passwords do not match')
    .run(req);

  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    req.flash("errors", errors.array().map(err => err.msg));
    return res.redirect('/profile');
  }
  try {
    const user = await User.findOne({ passwordResetToken: req.params.token })
      .where('passwordResetTokenExpires').gt(Date.now());
    if (!user) {
      req.flash('errors', 'Password reset token is invalid or has been expired');
      return res.redirect('/auth/forgot');
    }
    user.password = req.body.password;
    user.passwordResetToken = null;
    user.passwordResetTokenExpires = null;
    const savedUser = await user.save();
    req.logIn(savedUser, (loginErr) => {
      if (loginErr) {
        req.flash('error', 'Could not login');
        return res.redirect('/auth/login');
      }
      return res.redirect('/');
    });
  } catch (err) {
    next(err);
    return res.redirect('/auth/forgot');
  }
}
