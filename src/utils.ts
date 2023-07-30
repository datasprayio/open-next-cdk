// In cases where domainName consists of a Stack parameter or imported value,
// RecordSet construct fails to validate the name and erroneously suffixes with
// zone name. A workaround is to force a trailing dot to trick skipping the validation.
// https://github.com/aws/aws-cdk/issues/26572
export const domainAddTrailingDot = (domainName: string): string => {
  return domainName.endsWith('.') ? domainName : `${domainName}.`;
};
