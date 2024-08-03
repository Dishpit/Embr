# The Omega Programming Language

This is the main source code repository for Omega.

## Compiling the Interpreter

Quick and easy compilation:
`cl *.c /Fe:omega.exe`

Full compilation and deployment to production:
`cl *.c /Fe:omega.exe /O2 /Os /GL /DNDEBUG /MD /GS- /Gy /link /LTCG /OPT:REF /OPT:ICF /INCREMENTAL:NO /RELEASE`

## About Omega

Omega is (obviously) in its extreme infancy stage; it hasn't even attached to the uterus lining yet. Be advised that everything has the potential to change as time goes on, but I'm planning to prevent any breaking changes as much as possible so I can maintain a singular backwards-compatible spec as opposed to Rust's tendency to break previous versions.

Omega is designed to be a general-purpose programming language. Sure, it'd be great if I can set it up so it can accept multiple programming paradigms (functional, OOP, etc) in order to ease adoption from developers of all flavors, but that's wildly ambitious at this stage.

## Contributing
I'm thrilled that you're interested in contributing to Omega! Please adopt the following branch and development guidelines to ensure a smooth and collaborative contribution environment.

### Branch Guidelines

***

`main`

This is the primary branch. It contains the most stable and tested version of the language. All development eventually merges back into `main`, but only after thorough testing and review. The `main` branch is what the average user should be using in production.

***

`dev`

This branch serves as the primary integration branch for new features. It's the "working" version of the language and contains new features that are under development but might not be fully tested. Features should be developed in separate branches and merged into `dev` when they reach a stable state.

***

`feature/*`

For new features or significant changes, create individual branches off `dev`. Name these branches clearly based on the feature or change being implemented, e.g., `feature/new-syntax`, `feature/variable-bindings`. Once the feature is completed, tested, and approved, it gets merged back into the dev branch.

***

`release/*`

When `dev` reaches a state where it's ready to be released (after all features planned for the release are merged and tests are passed), create a release branch, e.g., `release/v0.1.0`. This branch is used for final testing and documentation updates. After it's fully tested and ready, it merges into `main` and back into `dev` to ensure any version updates or last-minute fixes are not lost.

***

`hotfix/*`

For urgent fixes that need to be applied directly to the stable version, create hotfix branches from `main`. For example, `hotfix/critical-bug-fix`. After the fix is implemented and tested, it merges back into both `main` and `dev` to ensure the fix is applied across all relevant versions.

***

### Development Guidelines

#### Develop in Isolation

Work on new features and non-urgent fixes in their respective `feature/*` branches to keep changes isolated until they're ready.

#### Regular Merges

Regularly merge changes from `main` into `dev` and then into feature branches to keep them up to date and minimize merge conflicts.

#### Code Reviews

Use pull requests for merging from `feature/*`, `release/*`, and `hotfix/*` branches to facilitate code reviews and ensure quality.
