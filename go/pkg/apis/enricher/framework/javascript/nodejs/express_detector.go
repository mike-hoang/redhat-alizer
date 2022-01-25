/*******************************************************************************
 * Copyright (c) 2021 Red Hat, Inc.
 * Distributed under license by Red Hat, Inc. All rights reserved.
 * This program is made available under the terms of the
 * Eclipse Public License v2.0 which accompanies this distribution,
 * and is available at http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 * Red Hat, Inc.
 ******************************************************************************/
package recognizer

import (
	"github.com/redhat-developer/alizer/pkg/apis/language"
)

type ExpressDetector struct{}

func (e ExpressDetector) DoFrameworkDetection(language *language.Language, config string) {
	if hasFramework(config, "express") {
		language.Frameworks = append(language.Frameworks, "Express")
	}
}
