package diags

// TODO: should be able to create a simpler/builder~esque api - diags.Env("VAR_NAME").Error("summary", "detail")

func NewError(component Component, path string, detail string, summary string) Diagnostic {
	return Diagnostic{
		Severity:  SeverityError,
		Component: component,
		Path:      path,
		Summary:   summary,
		Detail:    detail,
	}
}

func NewWarn(component Component, path string, detail string, summary string) Diagnostic {
	return Diagnostic{
		Severity:  SeverityWarn,
		Component: component,
		Path:      path,
		Summary:   summary,
		Detail:    detail,
	}
}

func NewInfo(component Component, path string, detail string, summary string) Diagnostic {
	return Diagnostic{
		Severity:  SeverityInfo,
		Component: component,
		Path:      path,
		Summary:   summary,
		Detail:    detail,
	}
}

func NewDebug(component Component, path string, detail string, summary string) Diagnostic {
	return Diagnostic{
		Severity:  SeverityDebug,
		Component: component,
		Path:      path,
		Summary:   summary,
		Detail:    detail,
	}
}
