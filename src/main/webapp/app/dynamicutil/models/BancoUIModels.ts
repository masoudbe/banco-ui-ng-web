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

export interface CommandDefinition {
  _iconSource: number;
  _title?: any;
  Name: string;
  Code: string;
  CommandType: number;
  QualifiedName: string;
  PresenterName?: any;
  DefaultFilterQualifiedName?: any;
  IconName?: any;
  ObjectDefinition?: any;
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
}

export interface CommandDefinition2 {
  _iconSource: number;
  _title?: any;
  Name: string;
  Code: string;
  CommandType: number;
  QualifiedName: string;
  PresenterName?: any;
  DefaultFilterQualifiedName?: any;
  IconName?: any;
  ObjectDefinition?: any;
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
}

export interface CommandDefinition3 {
  _iconSource: number;
  _title?: any;
  Name: string;
  Code: string;
  CommandType: number;
  QualifiedName: string;
  PresenterName?: any;
  DefaultFilterQualifiedName?: any;
  IconName?: any;
  ObjectDefinition?: any;
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
}

export interface CommandLinkHierarchy2 {
  Name: string;
  Order: number;
  CommandDefinition: CommandDefinition3;
  CommandLinkGroupId: number;
  CommandLinkHierarchies: any[];
  IsSelected: boolean;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
}

export interface ParentCommandLink {
  Name: string;
  Order: number;
  CommandDefinition?: any;
  CommandLinkGroupId: number;
  CommandLinkHierarchies: CommandLinkHierarchy2[];
  ParentCommandLink?: any;
  IsSelected: boolean;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
}

export interface CommandLinkHierarchy {
  Name: string;
  Order: number;
  CommandDefinition: CommandDefinition2;
  CommandLinkGroupId: number;
  CommandLinkHierarchies: any[];
  ParentCommandLink: ParentCommandLink;
  IsSelected: boolean;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
}

export interface CommandDefinition4 {
  _iconSource: number;
  _title?: any;
  Name: string;
  Code: string;
  CommandType: number;
  QualifiedName: string;
  PresenterName?: any;
  DefaultFilterQualifiedName?: any;
  IconName?: any;
  ObjectDefinition?: any;
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
}

export interface CommandLinkHierarchy3 {
  Name: string;
  Order: number;
  CommandDefinition: CommandDefinition4;
  CommandLinkGroupId: number;
  CommandLinkHierarchies: any[];
  IsSelected: boolean;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
}

export interface ParentCommandLink2 {
  Name: string;
  Order: number;
  CommandDefinition?: any;
  CommandLinkGroupId: number;
  CommandLinkHierarchies: CommandLinkHierarchy3[];
  ParentCommandLink?: any;
  IsSelected: boolean;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
}

export interface RootObject {
  Name: string;
  Order: number;
  CommandDefinition: CommandDefinition;
  CommandLinkGroupId: number;
  CommandLinkHierarchies: CommandLinkHierarchy[];
  ParentCommandLink: ParentCommandLink2;
  IsSelected: boolean;
  ID: number;
  Version: number;
  CreationDate: Date;
  LastChange: Date;
  Creator: string;
  LastUpdater: string;
}




