import { useGetSystemAppConfig } from '@/libs/hooks/useGetSystemAppConfig';
import { ReactElement } from 'react';

interface Props {
  children: ReactElement;
}

// Only renders children if the system is not Vydon Cloud
export default function OSSOnlyGuard(props: Props): ReactElement | null {
  const { children } = props;
  const { data: systemAppConfig } = useGetSystemAppConfig();

  if (systemAppConfig?.isVydonCloud) {
    return null;
  }
  return children;
}
