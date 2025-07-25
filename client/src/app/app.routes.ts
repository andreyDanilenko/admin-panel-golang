import { Routes } from '@angular/router';
import { AuthLayoutComponent } from './layouts/auth-layout/auth-layout.component';
import { NoAuthGuard } from './pages/guard/not-auth-guard';
import { AuthGuard } from './pages/guard/auth-guard';
import { AuthComponent } from './pages/auth/auth.component';
import { MainLayoutComponent } from './layouts/main-layout/main-layout.component';
import { HomeComponent } from './pages/main/main.component';
import { ChatsPageComponent } from './pages/chats/chats-page.component';
import { ChatPageComponent } from './pages/chat/chat-page.component';

export const routes: Routes = [
  {
    path: 'auth',
    component: AuthLayoutComponent,
    canActivate: [NoAuthGuard],
    children: [
      { path: 'login', component: AuthComponent },
      { path: '', redirectTo: 'login', pathMatch: 'full' }
    ]
  },
  {
    path: '',
    component: MainLayoutComponent,
    canActivate: [AuthGuard],
    children: [
      { path: '', component: HomeComponent },
    ]
  },
  {
    path: 'messages',
    component: MainLayoutComponent,
    canActivate: [AuthGuard],
    children: [
      { path: '', component: ChatsPageComponent },
      { path: ':id', component: ChatPageComponent },
    ]
  },
  { path: '**', redirectTo: '' }
];
