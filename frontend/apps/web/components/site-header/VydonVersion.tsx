'use client';
import { useQuery } from '@connectrpc/connect-query';
import { UserAccountService } from '@vydon/sdk';
import { ReactElement } from 'react';

export default function VydonVersion(): ReactElement | null {
  const { data } = useQuery(UserAccountService.method.getSystemInformation);
  if (!data?.version) {
    return null;
  }
  return <p className="text-sm tracking-tight">{data.version}</p>;
}
