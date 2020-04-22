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

export interface ObjectDefinition {
  Version: number;
  QualifiedName: string;
  DefaultView: string;
  DefaultList: string;
  IconName?: any;
  Name: string;
  Layouts: any[];
  ID: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
  Descriptor: string;
}

export interface CommandDefinition {
  Name: string;
  Code: string;
  CommandType: number;
  QualifiedName?: any;
  PresenterName: string;
  DefaultFilterQualifiedName: string;
  IconName?: any;
  IconSource: number;
  ObjectDefinition: ObjectDefinition;
  Title: string;
  ModuleId: number;
  DisplayState: number;
  TransactionCode?: any;
  IsParametric: boolean;
  ParentFilter?: any;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
  Descriptor: string;
}

export interface ObjectDefinition2 {
  Version: number;
  QualifiedName: string;
  DefaultView: string;
  DefaultList: string;
  IconName?: any;
  Name: string;
  Layouts: any[];
  ID: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
  Descriptor: string;
}

export interface CommandDefinition2 {
  Name: string;
  Code: string;
  CommandType: number;
  QualifiedName: string;
  PresenterName: string;
  DefaultFilterQualifiedName: string;
  IconName?: any;
  IconSource: number;
  ObjectDefinition: ObjectDefinition2;
  Title: string;
  ModuleId: number;
  DisplayState: number;
  TransactionCode?: any;
  IsParametric: boolean;
  ParentFilter?: any;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
  Descriptor: string;
}

export interface CommandLinkHierarchy2 {
  Name: string;
  Order: number;
  CommandDefinition: CommandDefinition2;
  CommandLinkGroupId: number;
  CommandLinkHierarchies: any[];
  Title: string;
  IsSelected: boolean;
  IsFolder: boolean;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
  Descriptor: string;
}

export interface CommandLinkHierarchy {
  Name: string;
  Order: number;
  CommandDefinition: CommandDefinition;
  CommandLinkGroupId: number;
  CommandLinkHierarchies: CommandLinkHierarchy2[];
  Title: string;
  IsSelected: boolean;
  IsFolder: boolean;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
  Descriptor: string;
}

export interface RootObject {
  Name: string;
  Order: number;
  CommandDefinition?: any;
  CommandLinkGroupId: number;
  CommandLinkHierarchies: CommandLinkHierarchy[];
  ParentCommandLink?: any;
  Title: string;
  IsSelected: boolean;
  IsFolder: boolean;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
  Descriptor: string;
}








