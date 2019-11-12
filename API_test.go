package dnsclub

import (
    "testing"
)

func TestGetLatestInfo(t *testing.T) {
    info, err := GetLatestInfo()
    if err != nil {
        t.Error(err)
    }

    if info.Code != 0 {
        t.Errorf("possible error in \"club.dns-shop.ru\". Expected code: 0, got %d", info.Code)
    }
}

func TestGetDigestPosts(t *testing.T) {
    info, err := GetDigestPosts()
    if err != nil {
        t.Error(err)
    }

    if info.Code != 0 {
        t.Errorf("possible error in \"club.dns-shop.ru\". Expected code: 0, got %d", info.Code)
    }

    if info.Data.Items[0].RootRubricType != "digest" {
        t.Errorf("The rubric does not match the function being called.\n Function rubric: digest\n Returned rubric: %s", info.Data.Items[0].RootRubricType)
    }
}

func TestGetBlogPosts(t *testing.T) {
    info, err := GetBlogPosts()
    if err != nil {
        t.Error(err)
    }

    if info.Code != 0 {
        t.Errorf("possible error in \"club.dns-shop.ru\". Expected code: 0, got %d", info.Code)
    }

    if info.Data.Items[0].RootRubricType != "blog" {
        t.Errorf("The rubric does not match the function being called.\n Function rubric: blog\n Returned rubric: %s", info.Data.Items[0].RootRubricType)
    }
}

func TestGetReviewPosts(t *testing.T) {
    info, err := GetReviewPosts()
    if err != nil {
        t.Error(err)
    }

    if info.Code != 0 {
        t.Errorf("possible error in \"club.dns-shop.ru\". Expected code: 0, got %d", info.Code)
    }

    if info.Data.Items[0].RootRubricType != "review" {
        t.Errorf("The rubric does not match the function being called.\n Function rubric: review\n Returned rubric: %s", info.Data.Items[0].RootRubricType)
    }
}