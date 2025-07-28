import { Catch, ArgumentsHost, ExceptionFilter } from '@nestjs/common';

@Catch()
export class AllExceptionsFilter implements ExceptionFilter {
  catch(exception: any, host: ArgumentsHost) {
    const ctx = host.switchToRpc();
    const response = ctx.getContext();

    // For gRPC, throw an error with status and message
    throw {
      status: 'error',
      message: exception.message || 'Internal server error',
    };
  }
}