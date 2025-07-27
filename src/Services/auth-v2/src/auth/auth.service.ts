import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { UsersService } from '../users/users.service';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    private jwtService: JwtService,
  ) {}

  async validateUser(email: string, password: string) {
    const user = await this.usersService.findByEmail(email);
    if (user && (await bcrypt.compare(password, user.password))) {
      const { password, ...result } = user;
      return result;
    }
    return null;
  }

  async login(email: string, password: string) {
    const user = await this.validateUser(email, password);
    if (!user) {
      throw new UnauthorizedException('Invalid credentials');
    }
    const payload = { email: user.email, sub: user.id };
    return {
      code: 0,
      accessToken: this.jwtService.sign(payload),
    };
  }

  async validateToken(token: string) {
    try {
      const decoded = this.jwtService.verify(token);

      console.log('Decoded token:', decoded);
      return {
        code: 0,
        user: {
          email: decoded.email,
          id: decoded.sub,
        }
      }
      // const user = await this.usersService.findByEmail(decoded.email);
      
      // if (!user) {
      //   throw new UnauthorizedException('User not found');
      // }
      
      // // Remove password from response
      // const { password, ...result } = user;
      
      // return {
      //   code: 0,
      //   email: result.email,
      //   id: result.sub,
      // };
    } catch (error) {
      return {
        code: 1,
        message: 'Invalid token'
      };
    }
  }
}