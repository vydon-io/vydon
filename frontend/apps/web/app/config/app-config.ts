export interface SystemAppConfig {
  isAuthEnabled: boolean;
  publicAppBaseUrl: string;
  posthog: PosthogConfig;
  unify: UnifyConfig;
  isVydonCloud: boolean;
  isStripeEnabled: boolean;
  enableRunLogs: boolean;
  signInProviderId?: string;
  isMetricsServiceEnabled: boolean;
  isJobHooksEnabled: boolean;
  isAccountHooksEnabled: boolean;
  isSlackAccountHookEnabled: boolean;
  gtag: GtagConfig;

  calendlyUpgradeLink: string;
  isGcpCloudStorageConnectionsEnabled: boolean;
  // server-side base url
  vydonApiBaseUrl: string;
  // public (client-side) base url;
  publicVydonApiBaseUrl: string;
  isRbacEnabled: boolean;
  isPiiDetectionJobEnabled: boolean;
}

interface PosthogConfig {
  enabled: boolean;
  key?: string;
  host: string;
}

interface UnifyConfig {
  enabled: boolean;
  key?: string;
}

interface GtagConfig {
  enabled: boolean;
  key?: string;
  conversion?: string;
}
