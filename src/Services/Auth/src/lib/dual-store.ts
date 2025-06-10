import { Store } from 'express-session'
import { RedisClientType } from 'redis'
import { SessionData } from 'express-session'
const MongoStore = require("connect-mongo");
import type { RedisStore } from "connect-redis"

export default class DualStore extends Store {
  constructor(
    private redisStore: RedisStore,
    private mongoStore: typeof MongoStore
  ) {
    super()
  }
  get = (sid: string, callback: (err: any, session?: SessionData | null) => void) => {
    this.redisStore.get(sid, (err: any, session: SessionData | null) => {
      if (err) return callback(err);
      if (session) return callback(null, session);
      this.mongoStore.get(sid, (merr: any, msession: SessionData | null) => {
        if (merr) {
          return callback(merr);
        }
        if (msession) {
          this.redisStore.set(sid, msession, (setErr) => {
            if (setErr) {
              console.error(`Redis cache-back failed ${setErr}`);
            }
          })
        }
        callback(null, msession);
      });
    });
  }

  set = (sid: string, session: SessionData, callback?: (err?: any) => void) => {
    let completed = 0;
    let isError = false;
    const done = (err?: any) => {
      if (err && !isError) {
        isError = true;
        callback?.(err);
      } else if (++completed == 2 && !isError) {
        callback?.();
      }
    }

    this.redisStore.set(sid, session, done);
    this.mongoStore.set(sid, session, done);
  }
  destroy = (sid: string, callback?: (err?: any) => void) => {
    let completed = 0;
    let isError = false;
    const done = (err?: any) => {
      if (err && !isError) {
        isError = true;
        callback?.(err);
      } else if (++completed == 2 && !isError) {
        callback?.();
      }
    }

    this.redisStore.destroy(sid, done);
    this.mongoStore.destroy(sid, done);
  }

  /** "Touches" a given session, resetting the idle timer. */
  touch = (sid: string, session: SessionData, callback?: () => void) => {
    let completed = 0;
    const done = () => {
      if (++completed == 2) {
        callback?.();
      }
    }

    if (this.redisStore.touch) {
      this.redisStore.touch(sid, session, done);
    } else {
      done();
    }

    if (this.mongoStore.touch) {
      this.mongoStore.touch(sid, session, done);
    } else {
      done();
    }
  }
}
