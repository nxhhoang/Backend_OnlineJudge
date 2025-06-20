import express from 'express';
import session from 'express-session'
import { config } from 'dotenv';
import mongoose from 'mongoose'
const MongoStore = require("connect-mongo");
import { v4 as uuid } from 'uuid'
import passport from 'passport'
import authRoutes from './routers/auth';
import debugRoutes from './routers/debug'
import authViewRoutes from './routers/auth.view'
import flash from "express-flash";
import lusca from 'lusca';
import { redisClient } from './redis/connection';
import { RedisStore } from "connect-redis"
import DualStore from './lib/dual-store';
import path from 'path';
import cors from 'cors';
import morgan = require('morgan');

config();

export default () => {
  let redisStore = new RedisStore({
    client: redisClient,
    ttl: 2, // 2 minutes
  });
  let mongoStore = MongoStore.create({
    client: mongoose.connection.getClient(),
    dbName: process.env.DATABASE_NAME,
    collectionName: 'session',
    stringify: false,
    autoRemove: 'interval',
    autoRemoveInterval: 2
  });
  const dualStore = new DualStore(redisStore, mongoStore);
  const app = express();

  app.use(express.json());
  app.use(express.urlencoded({ extended: true }));
  // app.set('views', path.join(__dirname, '../views'));
  // app.engine('html', require('ejs').renderFile); // Use EJS to render HTML
  // app.set('view engine', 'html');
  app.use(morgan(process.env.NODE_ENV === 'development' ? 'dev' : 'tiny'));
  app.use(
    session({
      genid: (_) => {
        return uuid();
      },
      secret: process.env.AUTH_SECRET || 'secret',
      resave: false,
      saveUninitialized: false,
      store: dualStore,
      cookie: {
        maxAge: 1000 * 60 * 60 * 24
      }
    })
  );
  app.use(cors({
    credentials: true
  }))
  app.use(passport.initialize());
  app.use(passport.session());
  app.use(flash());
  app.use(lusca({
    csrf: true,
    xframe: 'SAMEORIGIN',
    p3p: 'ABCDEF',
    hsts: { maxAge: 31536000, includeSubDomains: true, preload: true },
    xssProtection: true,
    nosniff: true,
    referrerPolicy: 'same-origin'
  }));

  app.use('/debug', debugRoutes);
  app.use('/auth', authRoutes);
  app.use('/view/auth', authViewRoutes);
  return app;
}
