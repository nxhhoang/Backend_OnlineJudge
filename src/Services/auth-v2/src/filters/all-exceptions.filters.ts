import { Catch, ArgumentsHost, ExceptionFilter } from '@nestjs/common';
import { Response } from 'express';

@Catch()
export class AllExceptionsFilter implements ExceptionFilter {
  catch(exception: any, host: ArgumentsHost) {
    const ctx = host.switchToRpc();
    const response = ctx.getContext<Response>();

    response.send({
      status: 'error',
      message: exception.message || 'Internal server error',
    });
  }
}