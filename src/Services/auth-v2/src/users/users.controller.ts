import { Controller, Post, Get, Body, Param, UseGuards } from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBody, ApiBearerAuth, ApiParam } from '@nestjs/swagger';
import { JwtAuthGuard } from '../auth/jwt.strategy';
import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-user.dto';
import { GetUser } from 'src/common/decorators';
import { JwtPayLoad } from 'src/common/model';

@ApiTags('Users')
@Controller('users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Post()
  @ApiOperation({ summary: 'Create a new user' })
  @ApiBody({ type: CreateUserDto })
  @ApiResponse({ 
    status: 201, 
    description: 'User created successfully',
    schema: {
      type: 'object',
      properties: {
        code: { type: 'number', example: 0 },
        user: {
          type: 'object',
          properties: {
            id: { type: 'number', example: 1 },
            username: { type: 'string', example: 'john_doe' },
            name: { type: 'string', example: 'John Doe' }
          }
        }
      }
    }
  })
  @ApiResponse({ status: 400, description: 'Bad request - validation failed' })
  @ApiResponse({ status: 409, description: 'User already exists' })
  async create(@Body() createUserDto: CreateUserDto) {
    return this.usersService.create(createUserDto);
  }

  @Get('email/:email')
  @ApiOperation({ summary: 'Find user by email' })
  @ApiBearerAuth('JWT-auth')
  @ApiParam({ name: 'email', description: 'User email address' })
  @ApiResponse({ 
    status: 200, 
    description: 'User found',
    schema: {
      type: 'object',
      properties: {
        id: { type: 'number', example: 1 },
        email: { type: 'string', example: 'user@example.com' },
        name: { type: 'string', example: 'John Doe' }
      }
    }
  })
  @ApiResponse({ status: 404, description: 'User not found' })
  async findByEmail(@Param('email') email: string) {
    console.log('Finding user by email:', email);
    return this.usersService.findByEmail(email);
  }

  @Get(':id')
  @ApiOperation({ summary: 'Get current authenticated user' })
  @ApiBearerAuth('JWT-auth')
  async getme(@Param('id') id: number) {
    const user = await this.usersService.findById(id);
    return {
      id: user.id,
      email: user.email,
      username: user.username,
      name: user.name,
    };
  }

  // @Get(':id')
  // @ApiOperation({ summary: 'Find user by ID' })
  // @ApiBearerAuth('JWT-auth')
  // @ApiParam({ name: 'id', description: 'User ID' })
  // @ApiResponse({ 
  //   status: 200, 
  //   description: 'User found',
  //   schema: {
  //     type: 'object',
  //     properties: {
  //       id: { type: 'number', example: 1 },
  //       email: { type: 'string', example: 'user@example.com' },
  //       name: { type: 'string', example: 'John Doe' }
  //     }
  //   }
  // })
  // @ApiResponse({ status: 404, description: 'User not found' })
  // async findById(@Param('id') id: string) {
  //   return this.usersService.findById(parseInt(id));
  // }
}