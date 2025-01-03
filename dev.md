https://github.com/ChromeGG/cel-js

	ve := personValidator.Validate(x)
	var nested proto.Message = x.GetHealth()
		if v, ok := nested.(protocheck.Validator); ok {
			ve = append(ve, v.Validate()...)
		}
	return ve


- remove panic in generated code