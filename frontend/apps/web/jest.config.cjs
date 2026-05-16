/** @type {import('jest').Config} */
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  verbose: true,
  moduleNameMapper: {
    '^@/(.*)$': '<rootDir>/$1',
  },
};
