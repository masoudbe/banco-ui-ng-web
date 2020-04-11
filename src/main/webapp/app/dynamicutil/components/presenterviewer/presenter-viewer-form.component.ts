import {Component, OnInit} from '@angular/core';
import {StoreService} from "app/dynamicutil/services/store.service";
import {FormControl} from "@angular/forms";

@Component({
  selector: 'bng-presenterviewer',
  templateUrl: './presenter-viewer-form.component.html',
  styleUrls: ['./presenter-viewer-form.component.scss']
})
export class PresenterViewerForm implements OnInit {

  tabs: string[] = [];
  selected = new FormControl(0);

  constructor(private storeService: StoreService) {
    this.storeService.presenters$
      .subscribe(val => this.addRemoveTab(val));
  }

  addTab(tabName: string): void {
    this.tabs.push(tabName);
    this.selected.setValue(this.tabs.length - 1);
  }

  removeTab(index: number): void {
    this.tabs.splice(index, 1);
  }

  private addRemoveTab(prsList: string[]): void {
    if (this.tabs.length < prsList.length) {
      this.addTab(prsList[prsList.length - 1]);
    } else {
      let rmTab = "";
      this.tabs.forEach(tab => {
        if (!prsList.includes(tab)) {
          rmTab = tab;
        }
      })

      if (rmTab !== "") {
        this.removeTab(this.tabs.indexOf(rmTab));
      }
    }
  }

  ngOnInit(): void {

  }

}
