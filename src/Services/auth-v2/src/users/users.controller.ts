import { Controller, UseGuards } from '@nestjs/common';
import { MessagePattern } from '@nestjs/microservices';
import { AuthGuard } from '@nestjs/passport';
import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-user.dto';

@Controller()
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @MessagePattern({ cmd: 'create_user' })
  async create(createUserDto: CreateUserDto) {
    return this.usersService.create(createUserDto);
  }

  @UseGuards(AuthGuard('jwt'))
  @MessagePattern({ cmd: 'find_user_by_email' })
  async findByEmail(email: string) {
    return this.usersService.findByEmail(email);
  }

  @UseGuards(AuthGuard('jwt'))
  @MessagePattern({ cmd: 'find_user_by_id' })
  async findById(id: string) {
    return this.usersService.findById(id);
  }
}