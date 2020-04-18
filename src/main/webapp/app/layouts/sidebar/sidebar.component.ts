import {Component, OnInit} from '@angular/core';
import {map, startWith} from "rxjs/operators";
import {noop, Observable} from "rxjs";
import {FormControl} from "@angular/forms";
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import * as CS from "app/dynamicutil/models/Constants";
import {StoreService} from "app/dynamicutil/services/store.service";
import {CommandDefinition, RootObject} from "app/dynamicutil/models/BancoUIModels";
import {isNull} from "app/shared/util/common-util";
import {JhiAlertService} from "ng-jhipster";

export interface SystemCode {
  ID: string;
  Name: string;
  Title: string;
}

@Component({
  selector: 'bng-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss']
})
export class SidebarComponent implements OnInit {

  filterTextControl = new FormControl();
  options: SystemCode[] = [];
  filteredOptions: Observable<SystemCode[]>;

  commandDefinitions: CommandDefinition[] = [];

  constructor(private dynamicService: DynamicService, private storeService: StoreService, private alertService: JhiAlertService) {

  }

  ngOnInit(): void {
    this.dynamicService.execute<any>(CS.GETSUBSYSTEMS, "")
      .pipe(
        map(data => {

          console.log("SUBSYSTEMSUBSYSTEMSUBSYSTEMSUBSYSTEM", data);
          const c = data;
          // const c = JSON.parse(data);
          console.log("SUBSYSTEMSUBSYSTEMSUBSYSTEMSUBSYSTEM22", c);
          return c.map((element: { [x: string]: any; }) => {
            const sc: SystemCode = {"ID": element["ID"], "Name": element["Name"], "Title": element["Title"]}
            return sc;
          });
        })
      )
      .subscribe(val => {
        this.options = val;
      }, err => console.log(err), noop);

    this.filteredOptions = this.filterTextControl.valueChanges
      .pipe(
        startWith(''),
        map(value => this._filter(value))
      );

    this.storeService.toggleSideBar$.subscribe()
  }

  private _filter(value: string): SystemCode[] {
    const filterValue = value.toString().toLowerCase();

    return this.options.filter(option => option.Name.toLowerCase().includes(filterValue));
  }

  subSystemSelected(systemId: string): void {

    this.dynamicService.execute<any>(CS.GETSUBSYSTEMCOMMANDS, "?CommandLinkGroupId=" + systemId)
      .pipe(
        map(data => {
            console.log("SYSTEMIDSYSTEMIDSYSTEMIDSYSTEMIDSYSTEMID", systemId);
            console.log("COMMANDSCOMMANDSCOMMANDSCOMMANDSCOMMANDS", data);
            const c = data;
            // const c = JSON.parse(data);
            console.log("COMMANDSCOMMANDSCOMMANDSCOMMANDSCOMMANDS22", c);
            const commands: CommandDefinition[] = [];
            for (let i = 0; i < c.length; i++) {
              if (!isNull(c[i].CommandDefinition)) {
                commands.push(c[i].CommandDefinition);
              }
              for (let j = 0; j < c[i].CommandLinkHierarchies.length; j++) {
                if (!isNull(c[i].CommandLinkHierarchies[j].CommandDefinition)) {
                  commands.push(c[i].CommandLinkHierarchies[j].CommandDefinition);
                }
              }
            }

            if(systemId === "4"){
              commands[1].Code = "30064";
              commands[1].Name = "AddCurrencyDailyRate";
            }

            return commands;
          }
        )
      )
      .subscribe(val => {
        this.commandDefinitions = val;
      }, noop, noop);
  }

  addPresenter(commandCode: string): void {
    this.storeService.addPresenter(commandCode);
  }
}
