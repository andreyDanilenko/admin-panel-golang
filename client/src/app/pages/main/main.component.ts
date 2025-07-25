import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './main.component.html',
  styles: [`
    h2 {
      color: #4a90e2;
      border-bottom: 2px solid #4a90e2;
      padding-bottom: 0.5rem;
      border-radius: 4px;
    }
    p {
      font-size: 1.1rem;
      color: #555;
    }
  `]
})
export class HomeComponent {}
