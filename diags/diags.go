// Package diags provides functionality to build a diagnostic understanding of
// user-supplied configuration values across multiple sources, including global
// config files, environment variables, local config files, and CLI flags.
//
// # Overview
//
// Diags aims to streamline the process of reporting errors, validation concerns
// and the underlying configuration values across various sources by providing
// a unified interface for collecting diagnostic information.
//
// By offering a consistent way to manage diagnostic information, diags helps
// end-users ensure that the application they are attempting to run is
// configured correctly. Diags helps to make configuration debuggable by
// understanding the exact source of where configuration is being overridden or
// what values are unexpected and why.
//
// By enabling the user to help themselves, the developer can provide a great
// onboarding experience and be assured that their application will run in a
// predictable manner due to being configured correctly!
//
// # Features
//
// The main features of the diags package include:
//
//   - Support for multiple configuration sources: Diags provides an API to
//     aggregate configuration values from global config files, environment
//     variables, local config files, and CLI flags.
//   - Comprehensive diagnostic reports: Diags can be used by developers to
//     generate detailed output that summarizes the configuration settings
//     across all sources, making it easier for end-users to understand how the
//     application is configured.
//
// # Conclusion
//
// Diags simplifies the process of user configuration of Go applications
// by providing a unified interface for developers to help diagnose and report
// configuration settings from multiple sources.
// By using diags, developers can ensure that applications are easier to
// configure, maintain and self error-correct for end-users.
package diags
