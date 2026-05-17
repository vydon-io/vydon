import type { NextRequest } from 'next/server';
import { GET as authGET, POST as authPOST } from './auth';

// next-auth 5.x ships its handlers with the legacy single-arg signature
// (req) => Promise<Response>, while Next 16 now expects the two-arg
// RouteHandlerConfig signature (request, context). Wrap to satisfy the
// new shape; the context object is unused by next-auth.
export async function GET(
  request: NextRequest,
  _context: { params: Promise<{ nextauth: string[] }> }
): Promise<Response> {
  return authGET(request);
}

export async function POST(
  request: NextRequest,
  _context: { params: Promise<{ nextauth: string[] }> }
): Promise<Response> {
  return authPOST(request);
}
