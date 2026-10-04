/**
 * STUDYHUB - Client-Side App Logic
 */

function openModal(modalId) {
  const modal = document.getElementById(modalId);
  if (modal) {
    modal.classList.add('open');
  }
}

function closeModal(modalId) {
  const modal = document.getElementById(modalId);
  if (modal) {
    modal.classList.remove('open');
  }
}

// Close modal when clicking on backdrop
document.addEventListener('click', function (e) {
  if (e.target.classList.contains('modal-backdrop')) {
    e.target.classList.remove('open');
  }
});

// Close modal on Escape key
document.addEventListener('keydown', function (e) {
  if (e.key === 'Escape') {
    document.querySelectorAll('.modal-backdrop.open').forEach(function (m) {
      m.classList.remove('open');
    });
  }
});

// Toast notification helper
function showToast(message, type) {
  let region = document.querySelector('.toast-region');
  if (!region) {
    region = document.createElement('div');
    region.className = 'toast-region';
    document.body.appendChild(region);
  }

  const toast = document.createElement('div');
  const alertClass = (type === 'success') ? 'alert-success' : 'alert-error';
  toast.className = 'alert ' + alertClass;

  const span = document.createElement('span');
  span.textContent = message;
  toast.appendChild(span);

  region.appendChild(toast);

  setTimeout(function () {
    toast.remove();
    if (region.children.length === 0) {
      region.remove();
    }
  }, 5000);
}

// HTMX Error Handling
document.addEventListener('htmx:beforeSwap', function (e) {
  const xhr = e.detail.xhr;
  const status = xhr ? xhr.status : 0;

  if (status >= 400) {
    const target = e.detail.target;
    if (target && (target.hasAttribute('data-swap-errors') || (target.closest && target.closest('[data-swap-errors]')))) {
      e.detail.shouldSwap = true;
      e.detail.isError = false;
    } else {
      e.detail.shouldSwap = false;
      let msg = '';
      const response = e.detail.serverResponse || (xhr && xhr.responseText) || '';
      if (response && response.trim()) {
        try {
          const parser = new DOMParser();
          const doc = parser.parseFromString(response, 'text/html');
          const alertEl = doc.querySelector('.alert');
          if (alertEl) {
            msg = alertEl.textContent.trim();
          }
        } catch (err) {
          // ignore parsing error
        }
      }

      if (!msg) {
        if (status === 403) {
          msg = "You don't have permission to do that.";
        } else if (status === 404) {
          msg = "Not found.";
        } else if (status >= 500) {
          msg = "Something went wrong. Please try again.";
        } else {
          msg = "Something went wrong. Please try again.";
        }
      }

      showToast(msg, 'error');
    }
  }
});

// Permission Pruning
function prunePermissions(root) {
  const scope = (root && root.querySelectorAll) ? root : document;
  const body = document.body;
  const viewerId = body.getAttribute('data-user-id');
  const viewerName = body.getAttribute('data-user-name');
  const viewerRole = body.getAttribute('data-user-role');
  if (!viewerRole) return;

  scope.querySelectorAll('[data-action]').forEach(function (el) {
    const action = el.getAttribute('data-action');
    let keep = true;

    if (action === 'post-edit') {
      const post = el.closest('[data-owner-name]');
      const isOwner = post && (post.getAttribute('data-owner-name') === viewerName);
      keep = (viewerRole === 'student') && isOwner;
    } else if (action === 'post-delete') {
      const post = el.closest('[data-owner-name]');
      const isOwner = post && (post.getAttribute('data-owner-name') === viewerName);
      keep = (viewerRole === 'student' && isOwner) || (viewerRole === 'admin');
    } else if (action === 'post-join') {
      keep = (viewerRole === 'student');
    } else if (action === 'course-delete') {
      const card = el.closest('[data-owner-id]');
      const isOwner = card && (String(card.getAttribute('data-owner-id')) === String(viewerId));
      keep = isOwner;
    }

    if (!keep) {
      el.remove();
    }
  });
}

// Unicode Initials Helper
function updateAvatars(root) {
  const scope = (root && root.querySelectorAll) ? root : document;
  scope.querySelectorAll('.avatar[data-initial]').forEach(function (el) {
    const name = el.getAttribute('data-initial');
    if (name && name.trim()) {
      const chars = Array.from(name.trim());
      if (chars.length > 0) {
        el.textContent = chars[0].toUpperCase();
      }
    }
  });
}

// Apply DOM updates on load and HTMX swap
function applyDomUpdates(root) {
  prunePermissions(root);
  updateAvatars(root);
}

if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', function () {
    applyDomUpdates(document);
  });
} else {
  applyDomUpdates(document);
}

document.addEventListener('htmx:afterSwap', function () {
  applyDomUpdates(document);
});

// File size limit check via data-max-mb
document.addEventListener('change', function (e) {
  const input = e.target;
  if (input && input.hasAttribute && input.hasAttribute('data-max-mb')) {
    const maxMb = parseFloat(input.getAttribute('data-max-mb'));
    const maxBytes = maxMb * 1024 * 1024;
    if (input.files && input.files[0]) {
      const file = input.files[0];
      if (file.size > maxBytes) {
        input.setCustomValidity('File size exceeds ' + maxMb + 'MB limit.');
        input.reportValidity();
      } else {
        input.setCustomValidity('');
      }
    } else {
      input.setCustomValidity('');
    }
  }
});

// Extract YouTube 11-char ID
function extractYouTubeID(url) {
  if (!url) return null;
  const regExp = /(?:youtube\.com\/(?:watch\?(?:.*&)?v=|embed\/|shorts\/|live\/)|youtu\.be\/)([A-Za-z0-9_-]{11})/;
  const match = url.match(regExp);
  return match ? match[1] : null;
}

// Admin course publish request preparation
function prepareCoursePublish(event) {
  const params = event.detail.parameters;
  const youtubeUrl = (params && params.youtube_url) ? params.youtube_url.trim() : '';
  const submissionId = params ? params.submission_id : '';

  const id = extractYouTubeID(youtubeUrl);
  const resultEl = document.getElementById('publish-result');

  if (!id) {
    event.preventDefault();
    if (resultEl) {
      resultEl.innerHTML = '<div class="alert alert-error"><span>Invalid YouTube URL. Please provide a valid YouTube link.</span></div>';
    }
    return;
  }

  if (resultEl) {
    resultEl.innerHTML = '';
  }

  event.detail.parameters.youtube_id = id;
  event.detail.path = '/course/' + encodeURIComponent(submissionId) + '/publish';
}

window.prepareCoursePublish = prepareCoursePublish;
window.showToast = showToast;
