import { Controller, UseGuards } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { AuthGuard } from '@nestjs/passport';
import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-user.dto';

@Controller()
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @GrpcMethod('AuthService', 'CreateUser')
  async create(data: CreateUserDto) {
    return this.usersService.create(data);
  }

  @UseGuards(AuthGuard('jwt'))
  @GrpcMethod('AuthService', 'FindUserByEmail')
  async findByEmail(data: { email: string; token: string }) {
    return this.usersService.findByEmail(data.email);
  }

  @UseGuards(AuthGuard('jwt'))
  @GrpcMethod('AuthService', 'FindUserById')
  async findById(data: { id: number; token: string }) { // Changed id to number
    return this.usersService.findById(data.id);
  }
}