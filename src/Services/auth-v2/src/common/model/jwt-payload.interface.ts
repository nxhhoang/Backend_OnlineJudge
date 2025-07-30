// import { UserRole } from "@prisma/client";

export interface JwtPayLoad {
  sub: number;
  email: string;
  iat?: number;
  exp?: number;
}
