import { Router } from "express";
import passport from 'passport';
import * as userControllers from '../controllers/user';
const router = Router();
import { config } from 'dotenv';
config();

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
    res.redirect((process.env.NODE_ENV === 'development' ? '/view/auth/home' : '/'));
  }
);

router.post('/login', userControllers.postLogin);
router.post('/signup', userControllers.postSignup);
router.get('/logout', userControllers.getLogout);
router.post('/forgot', userControllers.postForgot);
router.post('/password', userControllers.postPassword);
router.post('/reset/:token', userControllers.postReset);
router.get('/status', userControllers.getAuthStatus);
router.get('/csrftoken', userControllers.getCsrfToken);
export default router;
