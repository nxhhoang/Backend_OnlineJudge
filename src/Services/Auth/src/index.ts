import createApp from './app';
import { config } from 'dotenv';
import connectDB from './mongodb/connection';
import connectRedis from './redis/connection';
import { initPassport } from './config/passport';

config();

const init = async () => {
  try {
    await connectDB();
    await connectRedis();
    initPassport();
    const PORT = process.env.PORT || 3000;
    const app = createApp();
    app.listen(PORT, () => {
      console.log(`Running at ${PORT}`);
    });
  } catch (err) {
    throw err;
  }
}

init();
