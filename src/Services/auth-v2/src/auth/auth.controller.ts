import { Controller, Post, Body, Get, UseGuards, Request, Param } from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBody, ApiBearerAuth } from '@nestjs/swagger';
import { AuthService } from './auth.service';
import { JwtAuthGuard } from './jwt.strategy';
import { LoginDto } from './dto/login.dto';
import { RefreshTokenDto } from './dto/refresh-token.dto';

@ApiTags('Authentication')
@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('login')
  @ApiOperation({ summary: 'User login' })
  @ApiBody({ type: LoginDto })
  @ApiResponse({ 
    status: 200, 
    description: 'Login successful',
    schema: {
      type: 'object',
      properties: {
        code: { type: 'number', example: 0 },
        access_token: { type: 'string', example: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...' }
      }
    }
  })
  @ApiResponse({ status: 401, description: 'Invalid credentials' })
  async login(@Body() loginDto: LoginDto) {
    const access_token = await this.authService.login(loginDto.email, loginDto.password);
    return {
      message: "Login successful",
      access_token,
    }
  }

  @Get('validate/:token')
  @ApiOperation({ summary: 'Validate JWT token' })
  @ApiBearerAuth('JWT-auth')
  @ApiResponse({ 
    status: 200, 
    description: 'Token validation result',
    schema: {
      type: 'object',
      properties: {
        code: { type: 'number', example: 0 },
        user: {
          type: 'object',
          properties: {
            email: { type: 'string', example: 'user@example.com' },
            id: { type: 'number', example: 1 }
          }
        }
      }
    }
  })
  async validateToken(@Param('token') token: string) {
    return this.authService.validateToken(token);
  }

  @Post('refresh')
  @ApiOperation({ summary: 'Refresh JWT token' })
  @ApiBody({ type: RefreshTokenDto })
  @ApiResponse({ 
    status: 200, 
    description: 'Token refreshed successfully',
    schema: {
      type: 'object',
      properties: {
        code: { type: 'number', example: 0 },
        access_token: { type: 'string', example: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...' }
      }
    }
  })
  @ApiResponse({ status: 401, description: 'Invalid refresh token' })
  async refreshToken(@Body() refreshTokenDto: RefreshTokenDto) {
    const access_token = await this.authService.refreshToken(refreshTokenDto.refreshToken);
    return {
      message: "Token refreshed successfully",
      access_token: access_token,
    }
  }
}