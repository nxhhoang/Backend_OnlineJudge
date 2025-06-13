import passport from "passport";
import { User, IUser } from "../mongodb/user";
import { Strategy as GitHubStrategy } from 'passport-github2';
import { Strategy as LocalStrategy } from "passport-local";
import { Request, Response, NextFunction } from "express";
import { find } from "lodash";

const callbackURL = 'http://127.0.0.1:3000/auth/github/callback';

export function initPassport() {
  passport.use(
    new LocalStrategy({
      usernameField: "username",
      passwordField: "password"
    }, async (username, password, done) => {
      try {
        const user = await User.findOne({ username });
        if (!user) {
          return done(null, false, { message: "Username not found" });
        }
        user?.comparePassword(password, (err: Error | null, isMatch: boolean) => {
          if (err) {
            return done(err);
          }
          if (isMatch) {
            return done(null, user);
          }
          return done(null, false, { message: "Invalid username or password" })
        });
      } catch (err) {
        return done(err);
      }
    })
  );
  passport.use(
    new GitHubStrategy({
      clientID: process.env.AUTH_GITHUB_ID!,
      clientSecret: process.env.AUTH_GITHUB_SECRET!,
      callbackURL,
      passReqToCallback: true,
      scope: ['user:email', 'read:user']
    }, async (req: any, accessToken: string, refreshToken: string, profile: any, done: (error: any, user?: any) => void) => {
      let existingUser = await User.findOne({ githubId: profile.id });
      if (existingUser) {
        req.flash('errors', { msg: "There is already a Github account that belongs to you, Sign in with that account!" });
        return done(null, existingUser);
      }
      existingUser = await User.findOne({ email: profile.emails[0].value });
      if (existingUser) {
        req.flash('errors', 'A user with that email already exists.');
        return done(null, existingUser);
      }

      if (req.user) {
        try {
          const user = await User.findById(req.user.id);
          if (!user) {
            req.flash('errors', { msg: "There is no user with the given user id" });
            return done(null);
          }
          user.email = profile.emails[0].value ?? ''; // this shouldn't fail
          user.username = profile.username;
          user.githubId = profile.id;
          user.tokens.push({ kind: 'github', accessToken });
          user.profile.name = profile.displayName;
          user.profile.photos = profile.photos;
          const savedUser = await user.save();
          req.flash('success', { msg: "Github account has been linked" });
          done(null, savedUser);
        } catch (err) {
          if (err) {
            return done(err);
          }
        }
      } else {
        const user: any = new User({
          email: profile.emails[0].value ?? '', // this shouldn't fail
          githubId: profile.id,
          username: profile.username,
          profile: {
            name: profile.name,
            photos: profile.photos,
          },
          tokens: [{ kind: 'github', accessToken }]
        });
        const savedUser = await user.save();
        req.flash('success', { msg: "Github account has been linked" });
        done(null, savedUser);
      }
    })
  )
  passport.serializeUser((user: any, done) => {
    console.log(`serialize ${user.id}`);
    done(null, user.id);
  });

  passport.deserializeUser(async (id: string, done) => {
    console.log(`deserialize ${id}`);
    try {
      const user = await User.findById(id);
      done(null, user);
    } catch (err) {
      done(err);
    }
  });
}

// login middleware
export const isAuthenticated = (req: Request, res: Response, next: NextFunction) => {
  if (req.isAuthenticated()) {
    return next();
  }
  res.redirect('/auth/login');
}

// authorize middleware
export const isAuthorized = (req: Request, res: Response, next: NextFunction) => {
  const provider = req.path.split('/').slice(-1)[0];

  const user = req.user as IUser;
  if (find(user.tokens, { kind: provider })) {
    next();
  } else {
    res.redirect(`/auth/${provider}`);
  }
}
