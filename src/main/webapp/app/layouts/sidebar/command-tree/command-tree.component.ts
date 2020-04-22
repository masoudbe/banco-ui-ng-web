import {Component} from '@angular/core';
import {FlatTreeControl} from "@angular/cdk/tree";
import {MatTreeFlatDataSource, MatTreeFlattener} from "@angular/material/tree";
import * as CS from "app/dynamicutil/models/Constants";
import {map} from "rxjs/operators";
import {isNull} from "app/shared/util/common-util";
import {noop} from "rxjs";
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import {StoreService} from "app/dynamicutil/services/store.service";

interface CommandNode {
  name: string;
  title: string;
  code: string;
  qualifiedName: string;
  children?: CommandNode[];
}

interface FlatNode {
  expandable: boolean;
  name: string;
  level: number;
}

@Component({
  selector: 'bng-command-tree',
  templateUrl: './command-tree.component.html',
  styleUrls: ['./command-tree.component.scss']
})
export class CommandTreeComponent {

  private _transformer = (node: CommandNode, level: number) => {
    return {
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
    };
  }

  treeControl = new FlatTreeControl<FlatNode>(
    node => node.level, node => node.expandable);

  treeFlattener = new MatTreeFlattener(
    this._transformer, node => node.level, node => node.expandable, node => node.children);

  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  constructor(private dynamicService: DynamicService, private storeService: StoreService) {
  }

  hasChild = (_: number, node: FlatNode) => node.expandable;

  subSystemSelected(systemId: string): void {

    this.dynamicService.execute<any>(CS.GETSUBSYSTEMCOMMANDS, "?CommandLinkGroupId=" + systemId)
      .pipe(
        map(data => {
            console.log("SYSTEMIDSYSTEMIDSYSTEMIDSYSTEMIDSYSTEMID", systemId);
            console.log("COMMANDSCOMMANDSCOMMANDSCOMMANDSCOMMANDS", data);

            // HOMEDIFF
            // const c = data;
            const c = JSON.parse(data);
            console.log("COMMANDSCOMMANDSCOMMANDSCOMMANDSCOMMANDS22", c);

            const commandNodeArray: CommandNode[] = [];

            for (let i = 0; i < c.length; i++) {

              if (isNull(c[i].CommandDefinition)) {
                const cnParent: CommandNode = {code: c[i].ID, name: c[i].Name, title: c[i].Title , qualifiedName: '', children: []};
                commandNodeArray.push(cnParent);
                for (let j = 0; j < c[i].CommandLinkHierarchies.length; j++) {
                  const clh = c[i].CommandLinkHierarchies[j];
                  if (!isNull(clh.CommandDefinition)) {
                    const cnChild: CommandNode = {
                      code: clh.CommandDefinition.Code,
                      name: clh.CommandDefinition.Name,
                      title: clh.CommandDefinition.Title,
                      qualifiedName: clh.CommandDefinition.QualifiedName,
                      children: []
                    };
                    cnParent.children.push(cnChild);
                  }
                }
              } else {
                const cn: CommandNode = {
                  code: c[i].CommandDefinition.Code,
                  name: c[i].CommandDefinition.Name,
                  title: c[i].CommandDefinition.Title,
                  qualifiedName: c[i].CommandDefinition.QualifiedName,
                  children: []
                };
                commandNodeArray.push(cn);
              }
            }

            return commandNodeArray;
          }
        )
      )
      .subscribe(val => {
        console.log("CommandNodeCommandNodeCommandNodeCommandNode", val);
        this.dataSource.data = val;
      }, noop, noop);
  }

  executeCommand(node: any): void {
    const cn: CommandNode = this.findCommand(node.name, this.dataSource.data);
    this.storeService.addPresenter(cn.name);
  }

  findCommand(commandName: string, nodeArray: CommandNode[]): CommandNode {
    let result: CommandNode = undefined;
    for (let i = 0; i < nodeArray.length; i++) {
      const cn: CommandNode = nodeArray[i];
      if (cn.children.length === 0 && cn.name === commandName) {
        result = cn;
        return result;
      }
      result = this.findCommand(commandName, cn.children);
      if (!isNull(result)) {
        return result;
      }
    }

    return null;
  }
}
