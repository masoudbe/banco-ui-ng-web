import {Component, OnInit} from '@angular/core';
import {StoreService} from "app/dynamicutil/services/store.service";
import {FormControl} from "@angular/forms";
import {CommandInfo} from "app/dynamicutil/models/CommandInfo";
import {isNull} from "app/shared/util/common-util";

@Component({
  selector: 'bng-presenterviewer',
  templateUrl: './presenter-viewer-form.component.html',
  styleUrls: ['./presenter-viewer-form.component.scss']
})
export class PresenterViewerForm {

  tabs: CommandInfo[] = [];
  selected = new FormControl(0);

  constructor(private storeService: StoreService) {
    this.storeService.commands$
      .subscribe(val => this.addRemoveTab(val));
  }

  private addTab(ci: CommandInfo): void {
    this.tabs.push(ci);
    this.selected.setValue(this.tabs.length - 1);
  }

  removeTab(index: number): void {
    this.tabs.splice(index, 1);
  }

  private addRemoveTab(commandInfoList: CommandInfo[]): void {
    if (this.tabs.length < commandInfoList.length) {
      this.addTab(commandInfoList[commandInfoList.length - 1]);
    } else {
      let rmCi: CommandInfo = undefined;
      this.tabs.forEach(ci => {
        if (!commandInfoList.includes(ci)) {
          rmCi = ci;
        }
      })

      if (!isNull(rmCi)) {
        this.removeTab(this.tabs.indexOf(rmCi));
      }
    }
  }
}
