import {Component, OnInit, ViewChild} from '@angular/core';
import {map} from "rxjs/operators";
import {noop} from "rxjs";
import {DynamicService} from "app/dynamicutil/services/dynamic.service";
import * as CS from "app/dynamicutil/models/Constants";
import {StoreService} from "app/dynamicutil/services/store.service";
import {JhiAlertService} from "ng-jhipster";
import {CommandTreeComponent} from "app/layouts/sidebar/command-tree/command-tree.component";
import {CommandInfo} from "app/dynamicutil/models/CommandInfo";

export interface SystemData {
  ID: string;
  Name: string;
  Title: string;
  Image: string
}

@Component({
  selector: 'bng-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss']
})
export class SidebarComponent implements OnInit {

  loadSubSystems = false;
  step = -1;
  systemDataArray: SystemData[] = [];

  @ViewChild('commandTree') commandTree: CommandTreeComponent;

  constructor(private dynamicService: DynamicService, private storeService: StoreService, private alertService: JhiAlertService) {

  }

  ngOnInit(): void {
    this.loadSubSystems = false;
    this.dynamicService.execute<any>(CS.GETSUBSYSTEMS, "")
      .pipe(
        map(data => {

          console.log("SUBSYSTEMSUBSYSTEMSUBSYSTEMSUBSYSTEM", data);
          // HOMEDIFF
          // const c = data;
          const c = JSON.parse(data);
          console.log("SUBSYSTEMSUBSYSTEMSUBSYSTEMSUBSYSTEM22", c);

          return c.map((element: { [x: string]: any; }) => {
            const sc: SystemData = {
              "ID": element["ID"],
              Image: ''/*'data:image/png;base64,' + element["LinkGroupImage"]*/,
              "Name": element["Name"],
              "Title": element["Title"]
            }
            return sc;
          });
        })
      )
      .subscribe(val => {
        this.systemDataArray = val;
        console.log("111111111111111111111111111111SYSTEMSYSTEMSYSTEMSYSTEMSYSTEMSYSTEMSYSTEM", val);
        this.loadSubSystems = true;
      }, err => console.log(err), noop);

    this.storeService.toggleSideBar$.subscribe()
  }

  private addPresenter(commandCode: string): void {
    const ci: CommandInfo = {title: commandCode, code: commandCode, name: commandCode, qualifiedName: '', children: []};
    this.storeService.addPresenter(ci);
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

  addPresenterByCommand(commandTxt: HTMLInputElement, $event: any): void {
    if ($event.key === 'Enter') {
      this.addPresenter(commandTxt.value);
    }
  }
}
