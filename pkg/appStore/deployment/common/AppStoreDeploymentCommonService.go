/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package appStoreDeploymentCommon

import (
	appStoreBean "github.com/devtron-labs/devtron/pkg/appStore/bean"
	appStoreRepository "github.com/devtron-labs/devtron/pkg/appStore/repository"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
)

type AppStoreDeploymentCommonService interface {
	GetInstalledAppByClusterNamespaceAndName(clusterId int, namespace string, appName string) (*appStoreBean.InstallAppVersionDTO, error)
}

type AppStoreDeploymentCommonServiceImpl struct {
	logger                               *zap.SugaredLogger
	installedAppRepository               appStoreRepository.InstalledAppRepository
}

func NewAppStoreDeploymentCommonServiceImpl(logger *zap.SugaredLogger, installedAppRepository appStoreRepository.InstalledAppRepository) *AppStoreDeploymentCommonServiceImpl {
	return &AppStoreDeploymentCommonServiceImpl{
		logger:                               logger,
		installedAppRepository:               installedAppRepository,
	}
}

func (impl AppStoreDeploymentCommonServiceImpl) GetInstalledAppByClusterNamespaceAndName(clusterId int, namespace string, appName string) (*appStoreBean.InstallAppVersionDTO, error) {
	installedApps, err := impl.installedAppRepository.GetClusterComponentByClusterId(clusterId)

	if err != nil {
		if err == pg.ErrNoRows {
			impl.logger.Warnw("no installed apps found", "clusterId", clusterId)
			return nil, nil
		} else {
			impl.logger.Errorw("error while fetching installed apps", "clusterId", clusterId, "error", err)
			return nil, err
		}
	}

	for _, installedApp := range installedApps {
		if namespace == installedApp.Environment.Namespace && appName == installedApp.App.AppName {
			return impl.convert(installedApp), nil
		}
	}

	return nil, nil
}

//converts db object to bean
func (impl AppStoreDeploymentCommonServiceImpl) convert(chart *appStoreRepository.InstalledApps) *appStoreBean.InstallAppVersionDTO {
	return &appStoreBean.InstallAppVersionDTO{
		EnvironmentId:   chart.EnvironmentId,
		Id:              chart.Id,
		AppId:           chart.AppId,
		AppOfferingMode: chart.App.AppOfferingMode,
		ClusterId:       chart.Environment.ClusterId,
		Namespace:       chart.Environment.Namespace,
		AppName:         chart.App.AppName,
		EnvironmentName: chart.Environment.Name,
	}
}
