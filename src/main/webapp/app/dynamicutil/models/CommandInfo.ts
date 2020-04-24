export interface CommandInfo {
  name: string;
  title: string;
  code: string;
  qualifiedName: string;
  children?: CommandInfo[];
}
