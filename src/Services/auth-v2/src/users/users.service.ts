import { Injectable } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { CreateUserDto } from './dto/create-user.dto';
import * as bcrypt from 'bcrypt';

@Injectable()
export class UsersService {
  constructor(private prisma: PrismaService) {}

  async create(createUserDto: CreateUserDto) {
    const { username, password, name } = createUserDto;
    console.log('Creating user:', username, name);
    const hashedPassword = await bcrypt.hash(password, 10);
    return this.prisma.user.create({
      data: {
        username,
        password: hashedPassword,
        name,
      },
    });
  }

  async findByEmail(email: string) {
    return this.prisma.user.findUnique({ where: { email } });
  }

  async findById(user_id: number) { // Changed from string to number
    console.log('Finding user by user_id:', user_id);
    return this.prisma.user.findUnique({ where: { id: user_id } });
  }
}