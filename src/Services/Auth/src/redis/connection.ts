import { createClient } from 'redis'
require('dotenv').config();

let redisClient = createClient({
  url: process.env.REDIS_URI,
  password: process.env.REDIS_PASSWORD
});

export default async function connectRedis() {
  try {
    redisClient.on('error', (err) => console.log(err));
    await redisClient.connect();
    console.log("redis connection successful");
  } catch (err) {
    throw err;
  }
}

export { redisClient };
