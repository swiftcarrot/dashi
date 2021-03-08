import React from 'react';
import { FormLabel, FormControl, FormError } from '@swiftcarrot/react-form';

export default function FormGroup({ name, label, ...props }) {
  return (
    <div className="form-group">
      {label ? (
        <FormLabel name={name} className="form-label">
          {label}
        </FormLabel>
      ) : null}
      <FormControl name={name} className="form-control" {...props} />
      <FormError name={name} className="form-error" />
    </div>
  );
}
