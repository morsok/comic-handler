import { Component, input } from '@angular/core';
import { Comic } from '../services/comic.service';

@Component({
  selector: 'app-comic-card',
  template: `
    <div class="comic-card">
      <div class="comic-header">
        <h3 class="comic-title">{{ comic().title }}</h3>
        <span class="comic-id">#{{ comic().id }}</span>
      </div>
      <p class="comic-description">{{ comic().description }}</p>
      <div class="comic-actions">
        <button class="btn btn-primary">Read</button>
        <button class="btn btn-secondary">Details</button>
      </div>
    </div>
  `,
  styleUrl: './comic-card.component.scss'
})
export class ComicCardComponent {
  readonly comic = input.required<Comic>();
}
