;;; Package --- Summary
;;; Commentary:
;;
;; Set org config for publishing this website.

;;; Code:

(defconst html-style "<link rel=\"stylesheet\" type=\"text/css\" href=\"style.css\" />")

(defconst html-main-head
  (concat "<link rel=\"stylesheet\" type=\"text/css\" href=\"includes.css\" />" html-style))

(defconst html-posts-head
  (concat "<link rel=\"stylesheet\" type=\"text/css\" href=\"../includes.css\" />" html-style))

(defun postamble-text (text)
  "Wraps TEXT in a span with class postamble-text."
  (format "<span class=\"postamble-text\">%s</span>" text))

(defun postamble-version (version)
  "Wraps VERSION in a span with class version-number."
  (format "<span class=\"postamble-text version-number\">%s</span>" version))

(defun publish-nicolasknoebber-file ()
  "Exports current org file to html and uploads to s3://nicolasknoebber.com."
  (interactive)
  (org-publish-current-file)
  (let* (
	 (org-file (buffer-file-name (buffer-base-buffer)))
	 (publishing-dir (org-publish-property :publishing-directory
                                               (org-publish-get-project-from-filename org-file)))
         (html-file (replace-regexp-in-string "org$" "html" (buffer-name)))
	 (html-file-path (concat publishing-dir "/" html-file))
	 (site-path (replace-regexp-in-string ".+personal-website" "" html-file-path))
         (aws-s3-cmd
          (concat "aws s3 cp " html-file-path " s3://nicolasknoebber.com" site-path)))
    (eshell-command aws-s3-cmd)))

(defconst html-postamble
  (concat
   "<span id=\"made-with\">"
   (postamble-text "powered by&nbsp;&nbsp;")
   "<a href=\"https://www.gnu.org/software/emacs\">"
   "<img src=\"../logo/emacs.svg\" id=\"emacs-logo\" alt=\"Emacs\">"
   "</a>"
   (postamble-version emacs-version)
   "&nbsp<a href=\"https://orgmode.org\">"
   "<img src=\"../logo/org-mode.svg\" id=\"org-mode-logo\" alt=\"Org\">"
   "</a>"
   (postamble-version org-version)
   "</span>"
   "<span id=\"published\">"
   (format "%s" (format-time-string "%m/%e/%y"))
   "</span>"))

(defconst html-posts-postamble
  (concat
   html-postamble
   "<noscript><div id=\"no-script-comment-message\">Enable scripts to see and post comments.</div></noscript>"
   "<script type=\"text/javascript\" src=\"js/comments.js\"></script>"))

(setq org-publish-project-alist
      `(("personal-website"
         :components ("main" "posts"))
	("main"
	 :publishing-directory "~/projects/personal-website"
	 :base-directory "~/projects/personal-website/src"
	 :publishing-function org-html-publish-to-html
	 :section-numbers nil
	 :with-toc nil
	 :with-title nil
	 :html-head ,html-main-head
	 :html-preamble "<a href=\"/\">Nicolas Knoebber</a>"
	 :html-postamble ,html-postamble
	 :html-head-include-scripts nil
	 :html-head-include-default-style nil)
	("posts"
         :publishing-directory "~/projects/personal-website/posts"
	 :base-directory "~/projects/personal-website/src/posts"
	 :publishing-function org-html-publish-to-html
	 :section-numbers nil
	 :with-toc nil
	 :html-head ,html-posts-head
	 :html-head-include-scripts nil
	 :html-head-include-default-style nil
	 :html-preamble "<a href=\"../blog.html\">Blog</a>"
	 :html-postamble ,html-posts-postamble
	 )))


(add-to-list 'org-export-filter-timestamp-functions
             #'filter-timestamps)
(defun filter-timestamps(data backend _channel)
  "Remove <> around timestamps.  DATA is transcoded from the export BACKEND."
  (replace-regexp-in-string "&[lg]t;" "" data))

(setq-default org-display-custom-times t)
(setq org-time-stamp-custom-formats
      '("<%m/%d/%Y>" . "<%m/%d/%Y"))

;;; site.el ends here
