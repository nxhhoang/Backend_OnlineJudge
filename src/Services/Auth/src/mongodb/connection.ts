import mongoose from "mongoose";
require("dotenv").config();

export default async function connectDB() {
  if (!process.env.MONGO_URI) {
    throw new Error('Invalid/Missing environment variable: "MONGODB_URI"')
  }

  if (mongoose.connection.readyState == 0) {
    try {
      await mongoose.connect(
        process.env.MONGO_URI!,
        {
          maxPoolSize: 10,
        }
      );
      console.log("mongodb connection successful");
    } catch (err) {
      throw err;
    }
  }
};

