const { awscdk, JsonPatch, Tasks } = require('projen');

/**
 * To add additional platforms, ensure esbuild supports it here: https://esbuild.github.io/getting-started/#download-a-build
 */
const supportedOsCpus = [
  ['darwin', 'arm64'],
  ['darwin', 'x64'],
  ['linux', 'arm64'],
  ['linux', 'x64'],
  ['win32', 'arm64'],
  ['win32', 'x64'],
];
const supportedOses = [...new Set(supportedOsCpus.map(([os, _]) => os))];
const supportedCpus = [...new Set(supportedOsCpus.map(([_, cpu]) => cpu))];

const project = new awscdk.AwsCdkConstructLibrary({
  package: { installCommand: 'yarn install --ignoreAA-platform --check-files --frozen-lockfile' },
  author: 'Dataspray',
  authorAddress: 'matus@matus.io',
  cdkVersion: '2.115.0',
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
    gitBranch: 'go',
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
    // Bundle versions of esbuild for all supported platforms
    ...['esbuild', ...supportedOsCpus.map(([os, cpu]) => `@esbuild/${os}-${cpu}`)].map(
      (esBuildPackage) => `${esBuildPackage}@0.17.16` // If version changed, also chane in README.md
    ),
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

// Ignore platform check when installing esbuild for other supported platforms
// Replace in all tasks
[project.removeTask('install'), project.removeTask('install:ci')]
  .filter((task) => !!task)
  .forEach((task) => {
    project.addTask(task.name, {
      description: task.description,
      exec: task.steps[0].exec.replace('yarn install', 'yarn install --ignore-platform'),
    });
  });
// Replace in all Github Actions
// This is a very unstable way to patch the Github Actions workflows to add --ignore-platform.
project.github
  ?.tryFindWorkflow('build')
  ?.file?.patch(
    JsonPatch.replace('/jobs/build/steps/2/run', 'yarn install --ignore-platform --check-files'),
    JsonPatch.replace(
      '/jobs/package-js/steps/4/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    ),
    JsonPatch.replace(
      '/jobs/package-java/steps/5/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    ),
    JsonPatch.replace(
      '/jobs/package-python/steps/5/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    ),
    JsonPatch.replace(
      '/jobs/package-dotnet/steps/5/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    ),
    JsonPatch.replace(
      '/jobs/package-go/steps/5/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    )
  );
project.github
  ?.tryFindWorkflow('release')
  ?.file?.patch(
    JsonPatch.replace('/jobs/release/steps/3/run', 'yarn install --ignore-platform --check-files --frozen-lockfile'),
    JsonPatch.replace(
      '/jobs/release_npm/steps/4/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    ),
    JsonPatch.replace(
      '/jobs/release_maven/steps/5/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    ),
    JsonPatch.replace(
      '/jobs/release_pypi/steps/5/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    ),
    JsonPatch.replace(
      '/jobs/release_nuget/steps/5/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    ),
    JsonPatch.replace(
      '/jobs/release_golang/steps/5/run',
      'cd .repo && yarn install --ignore-platform --check-files --frozen-lockfile'
    )
  );
// An attempted more stable way is the following, but it doesn't work well with publishing actions particularly
// release_* jobs that are not rendered until syntch presumably
// const resolve = value => {
//   // Resolve based on projen/projen _resolve.ts
//   if (!value) return []
//   if (typeof value === 'function') return resolve(value.apply());
//   if (typeof value === 'object') return Object.entries(value);
//   if (value instanceof Map) return Array.from(value.values());
//   if (Array.isArray(value)) return value
//   return []
// }
// project.github.workflows.forEach(workflow => {
//   [...resolve(workflow.jobs), ...resolve(workflow.actions?.actions?.values())]
//     .forEach(([jobId, job]) => {
//       resolve(job.steps)?.forEach(([stepId, step]) => {
//         step.run = step.run?.replace('yarn install', 'yarn install --ignore-platform')
//       })
//     });
// });

// Patch package.json to adjust jest matching
const packageJson = project.tryFindObjectFile('package.json');
if (packageJson) {
  packageJson.patch(
    JsonPatch.replace('/jest/testMatch', [
      '<rootDir>/src/**/__tests__/**/*.ts?(x)',
      '<rootDir>/(test|src|assets)/**/*(*.)@(spec|test).ts?(x)',
    ])
  );
  // With esbuild, we only support these platforms
  packageJson.patch(JsonPatch.replace('/os', supportedOses));
  packageJson.patch(JsonPatch.replace('/cpu', supportedCpus));
}

project.synth();
