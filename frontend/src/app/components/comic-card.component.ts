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
  styles: [`
    .comic-card {
      background: var(--card-background);
      border: 1px solid var(--border-color);
      border-radius: 12px;
      padding: 1.5rem;
      transition: all 0.2s ease;
      box-shadow: 0 2px 8px var(--shadow-color);
    }

    .comic-card:hover {
      transform: translateY(-2px);
      box-shadow: 0 4px 16px var(--shadow-hover);
    }

    .comic-header {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      margin-bottom: 1rem;
    }

    .comic-title {
      font-size: 1.25rem;
      font-weight: 600;
      color: var(--text-primary);
      margin: 0;
      line-height: 1.3;
    }

    .comic-id {
      background: var(--accent-color);
      color: var(--accent-text);
      padding: 0.25rem 0.75rem;
      border-radius: 6px;
      font-size: 0.875rem;
      font-weight: 500;
    }

    .comic-description {
      color: var(--text-secondary);
      line-height: 1.5;
      margin-bottom: 1.5rem;
    }

    .comic-actions {
      display: flex;
      gap: 0.75rem;
    }

    .btn {
      padding: 0.5rem 1rem;
      border-radius: 6px;
      border: none;
      font-weight: 500;
      cursor: pointer;
      transition: all 0.2s ease;
    }

    .btn-primary {
      background: var(--primary-color);
      color: white;
    }

    .btn-primary:hover {
      background: var(--primary-hover);
    }

    .btn-secondary {
      background: transparent;
      color: var(--text-primary);
      border: 1px solid var(--border-color);
    }

    .btn-secondary:hover {
      background: var(--hover-background);
    }
  `]
})
export class ComicCardComponent {
  readonly comic = input.required<Comic>();
}
