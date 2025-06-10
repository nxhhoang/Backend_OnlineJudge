import crypto from 'crypto'
import nodemailer from 'nodemailer'
import passport from 'passport'
import { User, IUser, AuthToken } from '../mongodb/user'
import { Request, Response, NextFunction } from 'express'
import { IVerifyOptions } from 'passport-local'
import { WriteError } from 'mongodb'
import { Result, ValidationError, body, check, validationResult } from 'express-validator'
import '../config/passport'

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
    return res.redirect('/login');
  }

  passport.authenticate('local', (err: Error, user: IUser, info: IVerifyOptions) => {
    if (err) {
      req.flash('error', 'An error occurred during authentication');
      return res.redirect('/login');
    }
    if (!user) {
      req.flash('error', info.message || 'Invalid username or password');
      return res.redirect('/login');
    }

    req.logIn(user, (loginErr) => {
      if (loginErr) {
        req.flash('error', 'Could not login');
        return res.redirect('/login');
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
    res.redirect('/login');
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
    return res.redirect('/login');
  }

  const user = new User({
    email: req.body.email,
    password: req.body.password,
    username: req.body.username
  })

  try {
    if (await User.findOne({ email: req.body.email })) {
      req.flash('errors', 'A user with that email already exists.');
      return res.redirect('/signup');
    }

    if (await User.findOne({ username: req.body.username })) {
      req.flash('errors', 'A user with that username already exists.');
      return res.redirect('/signup');
    }
    const savedUser = await user.save();
    req.logIn(savedUser, (loginErr) => {
      if (loginErr) {
        req.flash('error', 'Could not login');
        return res.redirect('/login');
      }
      return res.redirect('/');
    });

  } catch (err) {
    req.flash('errors', 'An error occurred');
    return res.redirect('/signup');
  }
}
