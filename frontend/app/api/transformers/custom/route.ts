import { withNeosyncContext } from '@/api-only/neosync-context';
import {
  CreateCustomTransformerRequest,
  DeleteCustomTransformerRequest,
  GetCustomTransformersRequest,
  UpdateCustomTransformerRequest,
} from '@/neosync-api-client/mgmt/v1alpha1/transformer_pb';
import { NextRequest, NextResponse } from 'next/server';

export async function GET(req: NextRequest): Promise<NextResponse> {
  const { searchParams } = new URL(req.url);
  const accountId = searchParams.get('accountId') ?? '';
  return withNeosyncContext(async (ctx) => {
    return ctx.transformerClient.getCustomTransformers(
      new GetCustomTransformersRequest({ accountId })
    );
  })(req);
}

export async function DELETE(req: NextRequest): Promise<NextResponse> {
  const { searchParams } = new URL(req.url);
  const tId = searchParams.get('transformerId') ?? '';
  return withNeosyncContext(async (ctx) => {
    return ctx.transformerClient.deleteCustomTransformer(
      new DeleteCustomTransformerRequest({ transformerId: tId })
    );
  })(req);
}

export async function POST(req: NextRequest): Promise<NextResponse> {
  return withNeosyncContext(async (ctx) => {
    const body = CreateCustomTransformerRequest.fromJson(await req.json());
    return ctx.transformerClient.createCustomTransformer(body);
  })(req);
}

export async function PUT(req: NextRequest): Promise<NextResponse> {
  return withNeosyncContext(async (ctx) => {
    const body = UpdateCustomTransformerRequest.fromJson(await req.json());
    return ctx.transformerClient.updateCustomTransformer(body);
  })(req);
}
