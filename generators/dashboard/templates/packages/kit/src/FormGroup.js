import React from 'react';
import { FormLabel, FormControl, FormError } from '@swiftcarrot/react-form';

const FormGroup = ({ name, label, ...props }) => {
  return (
    <div className="form-group">
      <FormLabel name={name} className="form-label">
        {label || name}
      </FormLabel>
      <FormControl name={name} className="form-control" />
      <FormError name={name} className="form-error" />
    </div>
  );
};

export default FormGroup;
