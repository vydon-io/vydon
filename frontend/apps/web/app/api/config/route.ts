import { withVydonContext } from '@/api-only/vydon-context';
import { NextRequest, NextResponse } from 'next/server';
import { getSystemAppConfig } from './config';

export async function GET(req: NextRequest): Promise<NextResponse> {
  return withVydonContext(async () => getSystemAppConfig())(req);
}
