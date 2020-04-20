import {Component, OnInit, ViewChild} from '@angular/core';
import {map, startWith} from "rxjs/operators";
import {noop, Observable} from "rxjs";
import {FormControl} from "@angular/forms";
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import * as CS from "app/dynamicutil/models/Constants";
import {StoreService} from "app/dynamicutil/services/store.service";
import {CommandDefinition} from "app/dynamicutil/models/BancoUIModels";
import {isNull} from "app/shared/util/common-util";
import {JhiAlertService} from "ng-jhipster";
import {CommandTreeComponent} from "app/layouts/sidebar/command-tree/command-tree.component";

export interface SystemData {
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

  step = 0;
  systemDataArray: SystemData[] = [];

  @ViewChild('commandTree') commandTree: CommandTreeComponent;

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
            const sc: SystemData = {"ID": element["ID"], "Name": element["Name"], "Title": element["Title"]}
            return sc;
          });
        })
      )
      .subscribe(val => {
        this.systemDataArray = val;
      }, err => console.log(err), noop);

    this.storeService.toggleSideBar$.subscribe()
  }

  addPresenter(commandCode: string): void {
    this.storeService.addPresenter(commandCode);
  }

  setStep(index: number): void {
    this.step = index;
  }

  nextStep(): void {
    this.step++;
  }

  prevStep(): void {
    this.step--;
  }

  subSystemSelected(systemId: string, index: number): void {
    if (index !== this.step) {
      return;
    }

    this.commandTree.subSystemSelected(systemId);
  }
}
