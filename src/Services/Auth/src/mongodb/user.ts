import mongoose, { Schema, model, Document } from 'mongoose';
import bcrypt from 'bcrypt-nodejs';

export interface AuthToken {
  accessToken: string,
  kind: string
}

export interface IUser extends Document {
  githubId: string | null;
  googleId: string | null;

  email: string,
  username: string;
  password: string;
  passwordResetToken: String | null;
  passwordResetTokenExpires: Date | null;

  tokens: AuthToken[],
  profile: {
    name: string,
    gender: string,
    photos: { value: string }[],
  };

  comparePassword: comparePasswordFunction;
}

type comparePasswordFunction = (candidatePassword: string, done: (err: Error | null, isMatch: any) => void) => void;

export const UserSchema = new Schema<IUser>({
  githubId: String,
  googleId: String,

  email: { type: String, unique: true },
  username: { type: String, unique: true },
  password: String,
  passwordResetToken: String,
  passwordResetTokenExpires: Date,

  tokens: Array,
  profile: {
    name: String,
    gender: String,
    photos: [{ value: String }]
  },
}, {
  timestamps: true
});

UserSchema.pre('save', function(next) {
  const user = this as IUser;
  if (!user.isModified('password')) {
    return next();
  }
  bcrypt.genSalt(10, (err, salt) => {
    if (err) {
      return next(err);
    }
    bcrypt.hash(user.password, salt, null, (err: mongoose.Error, hash) => {
      if (err) {
        return next(err);
      }
      user.password = hash;
      next();
    });
  });
});

UserSchema.methods.comparePassword = function(candidatePassword: any, done: (err: Error | null, isMatch: boolean) => void) {
  bcrypt.compare(candidatePassword, this.password, function(err: Error | null, isMatch: boolean) {
    done(err, isMatch);
  });
};


export const User = model<IUser>('User', UserSchema);
