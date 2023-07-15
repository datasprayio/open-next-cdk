const { awscdk, JsonPatch } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Dataspray',
  authorAddress: 'matus@matus.io',
  cdkVersion: '2.73.0',
  defaultReleaseBranch: 'main',
  dependabot: true,
  name: 'OpenNext via CDK',
  repositoryUrl: 'https://github.com/datasprayio/open-next-cdk.git',
  authorOrganization: true,
  packageName: 'open-next-cdk',
  publishToMaven: {
    javaPackage: 'io.dataspray.opennextcdk',
    mavenGroupId: 'io.dataspray',
    mavenArtifactId: 'open-next-cdk',
    mavenEndpoint: 'https://s01.oss.sonatype.org',
  },
  publishToPypi: {
    distName: 'open-next-cdk',
    module: 'open_next_cdk',
  },
  publishToGo: {
    moduleName: 'github.com/datasprayio/open-next-cdk',
  },
  publishToNuget: {
    packageId: 'Dataspray.OpenNextCdk',
    dotNetNamespace: 'Dataspray.OpenNextCdk',
  },
  description: 'Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK',
  keywords: ['open-next', 'opennext', 'nextjs', 'next', 'aws-cdk', 'awscdk', 'cdk', 'serverless'],
  eslintOptions: {
    prettier: true,
  },

  majorVersion: 0,

  tsconfig: { compilerOptions: { noUnusedLocals: false }, include: ['assets/**/*.ts'] },
  tsconfigDev: { compilerOptions: { noUnusedLocals: false } },

  bundledDeps: [
    'cross-spawn',
    'fs-extra',
    'indent-string',
    'micromatch',
    '@aws-sdk/client-s3',
    '@types/cross-spawn',
    '@types/fs-extra',
    '@types/micromatch',
    '@types/aws-lambda',
    'esbuild@0.17.16',
    'aws-lambda',
    'serverless-http',
    'jszip',
    'glob',
    'node-fetch',
    '@aws-sdk/signature-v4',
    '@aws-crypto/sha256-js',
  ] /* Runtime dependencies of this module. */,
  devDeps: ['open-next', 'aws-sdk', 'constructs'] /* Build dependencies for this module. */,

  // do not generate sample test files
  sampleCode: false,
});
const packageJson = project.tryFindObjectFile('package.json');
if (packageJson) {
  packageJson.patch(
    JsonPatch.replace('/jest/testMatch', [
      '<rootDir>/src/**/__tests__/**/*.ts?(x)',
      '<rootDir>/(test|src|assets)/**/*(*.)@(spec|test).ts?(x)',
    ])
  );
}

project.synth();
