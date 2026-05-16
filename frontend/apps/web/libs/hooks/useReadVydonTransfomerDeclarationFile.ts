import { useQuery, UseQueryResult } from '@tanstack/react-query';
import { fetcher } from '../fetcher';

export function useReadVydonTransformerDeclarationFile(): UseQueryResult<string> {
  return useQuery({
    queryKey: [`/api/files/vydon-transformer-declarations`],
    queryFn: (ctx) => fetcher(ctx.queryKey.join('/')),
  });
}
