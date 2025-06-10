import { Router } from "express";
import passport from 'passport';
import { Request, Response, NextFunction } from "express";

const router = Router();

router.get(
  '/github',
  passport.authenticate(
    'github',
    { scope: ['user:email'] }
  )
);

router.get(
  '/github/callback',
  passport.authenticate('github', { failureRedirect: '/login' }),
  (req, res) => {
    res.redirect('/profile');
  }
);


export default router;
