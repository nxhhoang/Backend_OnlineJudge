import { Request, Response, NextFunction } from "express";
import { Router } from "express";
const MongoStore = require("connect-mongo");
import { redisClient } from '../redis/connection';
import { RedisStore } from "connect-redis"
import mongoose from "mongoose";

const router = Router();

router.get('/api/v1/debug-session', (req: Request, res: Response) => {
  res.json({
    sessionID: req.sessionID,
    user: req.user,
    session: req.session
  });
});

router.get('/api/v1/debug-session/:id', async (req, res) => {
  const sessionId = req.params.id;
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
  // Check Redis
  redisStore.get(sessionId, (redisErr, redisSession) => {
    // Check MongoDB
    mongoStore.get(sessionId, (mongoErr: any, mongoSession: any) => {
      res.json({
        sessionId,
        redis: {
          error: redisErr,
          session: redisSession
        },
        mongo: {
          error: mongoErr,
          session: mongoSession
        }
      });
    });
  });
});

export default router;
