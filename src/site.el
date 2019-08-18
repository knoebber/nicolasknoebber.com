;;; Package --- Summary
;;; Commentary:
;;
;; Set org config for publishing this website.
;;
;; Site Map
;; /index.html
;; /index.org
;; /posts/index.hmtl
;; /posts/[name].html
;; /posts/[name].html
;; /posts/org/[name].org
;; /posts/images/

;;; Code:
(setq org-publish-project-alist
      '(("personal-website"
	 :base-directory "~/projects/personal-website/src"
	 :publishing-directory "~/projects/personal-website"
         :publishing-function org-html-publish-to-html
         :recursive t
	 ;; :style "<link rel=\"stylesheet\"
	 ;; 	href=\"style.css\"
	 ;; 	type=\"text/css\"/>"
	 )))
;;; site.el ends here
