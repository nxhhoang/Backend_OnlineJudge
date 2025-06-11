import { Router } from "express";
import passport from 'passport';
import { Request, Response, NextFunction } from "express";
import * as userControllers from '../controllers/user';

const router = Router();

router.get(
  '/github',
  passport.authenticate(
    'github',
    { scope: ['user:email', 'read:user'] }
  )
);

router.get(
  '/github/callback',
  passport.authenticate('github', { failureRedirect: '/auth/login' }),
  (req, res) => {
    res.redirect('/profile');
  }
);

router.get('/login', userControllers.getLogin);
router.post('/login', userControllers.postLogin);
router.get('/signup', userControllers.getSignup);
router.post('/signup', userControllers.postSignup);
router.get('/logout', userControllers.getLogout);
router.get('/forgot', userControllers.getForgot);
router.post('/forgot', userControllers.postForgot);
router.get('/reset/:token', userControllers.getReset);
router.post('/reset/:token', userControllers.postReset);

export default router;
