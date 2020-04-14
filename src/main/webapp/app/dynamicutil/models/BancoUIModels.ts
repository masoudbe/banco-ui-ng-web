export interface ControlItem {
  IsEnabled: boolean;
  IsVisible: boolean;
  IsRequired: boolean;
  Name: string;
  Caption: string;
  ControlType: number;
  Container?: any;
  Order: number;
  ColumnSpan: number;
}

export interface Model {
}

export interface UICommand {
  Title: string;
  WebExecuter: boolean;
  Name: string;
  Description: string;
  Image?: any;
}

export interface TemplateOptions {
  label: string;
  required: boolean;
  type: string;
}

export interface FieldGroup {
  key: string;
  className: string;
  type: string;
  templateOptions: TemplateOptions;
}

export interface Field {
  fieldGroupClassName: string;
  fieldGroup: FieldGroup[];
}

export interface Tab {
  label: string;
  fields: Field[];
}


