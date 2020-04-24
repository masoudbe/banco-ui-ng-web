import { Route } from '@angular/router';
import {ToolbarComponent} from "app/layouts/toolbar/toolbar.component";

export const toolbarRoute: Route = {
  path: '',
  component: ToolbarComponent,
  outlet: 'toolbar'
};
