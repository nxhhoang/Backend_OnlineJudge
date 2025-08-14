// import { UserRole } from "@prisma/client";

export interface JwtPayLoad {
  sub: number;
  username: string;
  iat?: number;
  exp?: number;
}
